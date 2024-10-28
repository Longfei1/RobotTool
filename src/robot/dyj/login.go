package dyj

import (
	"RobotTool/src/common"
	"RobotTool/src/jsdef/config"
	"RobotTool/src/robot"
	"RobotTool/src/robot/dyj/pbgo"
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	log "github.com/sirupsen/logrus"
	"time"
)

type Login struct {
	core.Action
}

func (this *Login) OnTick(tick *core.Tick) b3.Status {
	rob := tick.Blackboard.GetMem(common.Robot).(robot.IRobot)
	serverCfg := tick.Blackboard.GetMem(common.ServerConfig).(*config.ServerConfig)

	_ = rob.Close()
	err := rob.Connect(serverCfg.Addr)
	if err != nil {
		common.LogError(tick, err)
		return b3.ERROR
	}

	loginReq := &pbgo.C2SLogin{
		Timestamp: time.Now().Unix(),
		UserId:    serverCfg.UserId,
	}

	err = rob.SendMsg(loginReq)
	if err != nil {
		common.LogError(tick, err)
		return b3.ERROR
	}

	log.Info("login success")
	return b3.SUCCESS
}
