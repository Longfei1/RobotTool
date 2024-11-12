package app

import (
	"RobotTool/src/common"
	"RobotTool/src/jsdef/config"
	"RobotTool/src/jsdef/jsmsg"
	"RobotTool/src/robot"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Longfei1/lorca"
	"github.com/gin-gonic/gin"
	b3 "github.com/magicsea/behavior3go"
	bevCfg "github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type App struct {
	ui  lorca.UI
	web *gin.Engine

	robot        robot.IRobot
	btProjectCfg *bevCfg.BTProjectCfg
	serverCfg    *config.ServerConfig
}

func NewApp(r robot.IRobot) *App {
	return &App{
		robot: r,
	}
}

func (a *App) Run() {
	if err := a.initBev(); err != nil {
		log.Error(err)
		return
	}

	if err := a.initWeb(); err != nil {
		log.Error(err)
		return
	}

	if err := a.initUi(); err != nil {
		log.Error(err)
		return
	}

	a.robot.SetAppApi(a) //注册回调

	log.Info("app run ...")

	// 等待 UI 窗口关闭
	defer a.ui.Close()
	<-a.ui.Done()

	_ = a.robot.Close()
}

func (a *App) initWeb() error {
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	a.web = gin.Default()

	a.web.StaticFS("/", gin.Dir(path.Join(workDir, "ui/dist"), true))
	go func() {
		err := a.web.Run(":8888")
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (a *App) initUi() error {
	//创建缓存目录
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	cacheDir := path.Join(workDir, "cache/chrome")
	if err = os.MkdirAll(cacheDir, 0755); err != nil {
		return err
	}

	//ui, err := lorca.New(path.Join(workDir, "ui/dist/index.html"),
	ui, err := lorca.New("http://localhost:8888",
		cacheDir, 1024, 768)
	if err != nil {
		return err
	}

	// 绑定 Go 函数到 JS 的 'go' 对象
	_ = ui.Bind("hello", func() string {
		return "Hello from Go!"
	})
	_ = ui.Bind("execBevTree", a.execBevTree)
	_ = ui.Bind("initServerConfig", a.initServerConfig)
	_ = ui.Bind("getProtoRequest", a.getProtoRequest)
	_ = ui.Bind("executeRequest", a.executeRequest)

	a.ui = ui

	return nil
}

func (a *App) initBev() error {
	projectCfg, ok := bevCfg.LoadProjectCfg("ui/public/conf/behavior.json")
	if !ok {
		return errors.New("load behavior config failed")
	}
	a.btProjectCfg = projectCfg

	return nil
}

func (a *App) execBevTree(id string) error {
	//自定义节点注册
	maps := b3.NewRegisterStructMaps()
	a.robot.RegisterBevMap(maps)

	var tree *core.BehaviorTree
	//载入
	for _, v := range a.btProjectCfg.Trees {
		if v.ID != id {
			continue
		}
		t := loader.CreateBevTreeFromConfig(&v, maps)
		tree = t
	}

	if tree == nil {
		return fmt.Errorf("not found tree id:%v", id)
	}

	//输入板
	board := core.NewBlackboard()
	board.SetMem(common.Robot, a.robot)
	board.SetMem(common.ServerConfig, a.serverCfg)

	go func() {
		var cnt int
		for {
			ret := tree.Tick(cnt, board)
			if ret == b3.RUNNING {
				cnt++
				time.Sleep(time.Millisecond * 100)
				continue
			}

			log.Infof("bev tree ret:%v", ret)
			if ret != b3.SUCCESS {
				err, ok := board.GetMem(common.ErrorStr).(error)
				if ok {
					a.ShowPopMsg(fmt.Sprintf("行为树：%v 运行失败\n错误：%v", tree.GetTitile(), err.Error()))
				} else {
					a.ShowPopMsg(fmt.Sprintf("行为树：%v 运行失败", tree.GetTitile()))
				}
			}

			break
		}
	}()
	return nil
}

func (a *App) initServerConfig(cfg *config.ServerConfig) error {
	a.serverCfg = cfg
	return nil
}

func (a *App) getProtoRequest() []*jsmsg.ProtoMsg {
	return a.robot.GetProtoRequest()
}

func (a *App) executeRequest(req *jsmsg.RequestMsgData) error {
	return a.robot.SendRequestMsg(req)
}

// **** js函数 ****//
func (a *App) AddShowMsg(msg *jsmsg.ShowMessage) {
	arg, err := json.Marshal(msg)
	if err != nil {
		log.Error(err)
		return
	}

	_ = a.ui.Eval(fmt.Sprintf("addShowMsg('%s')", string(arg)))
}

func (a *App) AddTipTag(tag *jsmsg.TipTag) {
	arg, err := json.Marshal(tag)
	if err != nil {
		log.Error(err)
		return
	}

	_ = a.ui.Eval(fmt.Sprintf("addTipTag('%s')", string(arg)))
}

func (a *App) AddTipMsg(msg string) {
	_ = a.ui.Eval(fmt.Sprintf("addTipMsg('%s')", msg))
}

func (a *App) ShowPopMsg(msg string) {
	_ = a.ui.Eval(fmt.Sprintf("showPopMsg('%s')", msg))
}

// **** js函数 ****//
