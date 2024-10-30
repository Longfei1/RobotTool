package robot

import (
	"RobotTool/src/app/api"
	"RobotTool/src/jsdef/jsmsg"
	"github.com/magicsea/behavior3go"
)

type INetWork interface {
	Connect(url string) error
	Close() error
	SendData(data []byte) error
	SendMsg(args ...interface{}) error
}

type IRobot interface {
	INetWork

	RegisterBevMap(mp *behavior3go.RegisterStructMaps)
	SetAppApi(appApi api.IAppApi)
	GetProtoRequest() []*jsmsg.ProtoMsg
	SendRequestMsg(req *jsmsg.RequestMsgData) error
}
