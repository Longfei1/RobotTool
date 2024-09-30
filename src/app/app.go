package app

import (
	"RobotTool/src/robot"
	"errors"
	"fmt"
	"github.com/Longfei1/lorca"
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
	log "github.com/sirupsen/logrus"
)

type App struct {
	ui lorca.UI

	robot        robot.IRobot
	btProjectCfg *config.BTProjectCfg
}

func NewApp(r robot.IRobot) *App {
	return &App{
		robot: r,
	}
}

func (a *App) Run() {
	if err := a.initUi(); err != nil {
		log.Error(err)
		return
	}

	if err := a.initBev(); err != nil {
		log.Error(err)
		return
	}

	// 等待 UI 窗口关闭
	defer a.ui.Close()
	<-a.ui.Done()
}

func (a *App) initUi() error {
	ui, err := lorca.New("http://localhost:5173",
		"", 1024, 768, "--remote-allow-origins=*")
	if err != nil {
		return err
	}

	// 绑定 Go 函数到 JS 的 'go' 对象
	_ = ui.Bind("hello", func() string {
		return "Hello from Go!"
	})

	a.ui = ui

	return nil
}

func (a *App) initBev() error {
	projectCfg, ok := config.LoadProjectCfg("ui/public/conf/behavior.b3")
	if !ok {
		return errors.New("load behavior config failed")
	}
	a.btProjectCfg = projectCfg

	return nil
}

func (a *App) ExecBevTree(id string) error {
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
	//循环每一帧
	for i := 0; i < 5; i++ {
		_ = tree.Tick(i, board)
	}
	return nil
}
