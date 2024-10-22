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
	b3 "github.com/magicsea/behavior3go"
	bevCfg "github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
	log "github.com/sirupsen/logrus"
	"time"
)

type App struct {
	ui lorca.UI

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

func (a *App) initUi() error {
	ui, err := lorca.New("http://localhost:5173",
		"", 1024, 768)
	if err != nil {
		return err
	}

	// 绑定 Go 函数到 JS 的 'go' 对象
	_ = ui.Bind("hello", func() string {
		return "Hello from Go!"
	})
	_ = ui.Bind("execBevTree", a.execBevTree)
	_ = ui.Bind("initServerConfig", a.initServerConfig)

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
			break
		}
	}()
	return nil
}

func (a *App) initServerConfig(cfg *config.ServerConfig) error {
	a.serverCfg = cfg
	return nil
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

// **** js函数 ****//
