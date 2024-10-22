package robot

import "github.com/magicsea/behavior3go"

type INetWork interface {
	Connect(url string) error
	Close() error
	SendData(data []byte) error
	SendMsg(args ...interface{}) error
}

type IRobotCallback interface {
	OnMessage(msg interface{})
}

type IRobot interface {
	INetWork

	RegisterBevMap(mp *behavior3go.RegisterStructMaps)
	SetCallback(callback IRobotCallback)
}
