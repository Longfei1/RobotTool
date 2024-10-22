package dyj

import (
	"RobotTool/src/network"
	"RobotTool/src/robot"
	"RobotTool/src/robot/dyj/pbgo"
	"errors"
	"github.com/magicsea/behavior3go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"time"
)

type RobotDyj struct {
	wsClient     *network.WsClient
	msgProcessor *MsgProcessor
	callback     robot.IRobotCallback
}

func NewRobotDyj() *RobotDyj {
	rob := &RobotDyj{
		msgProcessor: NewMsgProcessor(),
	}

	registerWsMsg(rob.msgProcessor)

	return rob
}

func (r *RobotDyj) RegisterBevMap(mp *behavior3go.RegisterStructMaps) {
	mp.Register("Login", &Login{})
}

func (r *RobotDyj) SetCallback(callback robot.IRobotCallback) {
	r.callback = callback
}

func (r *RobotDyj) Connect(url string) error {
	if r.wsClient != nil {
		return errors.New("already connect")
	}

	client := network.NewWsClient(r)
	err := client.Connect(url)
	if err != nil {
		return err
	}

	r.wsClient = client

	return nil
}

func (r *RobotDyj) Close() error {
	if r.wsClient == nil {
		return nil
	}

	_ = r.wsClient.Close()
	return nil
}

func (r *RobotDyj) SendMsg(args ...interface{}) error {
	//参数校验
	if len(args) < 1 {
		return errors.New("args len < 1")
	}
	msg, ok := args[0].(proto.Message)
	if !ok {
		return errors.New("arg[0] not proto.Message")
	}

	log.Debugf("send msg: %v", msg)

	data, err := r.msgProcessor.Marshal(msg)
	if err != nil {
		return err
	}
	return r.SendData(data)
}

func (r *RobotDyj) SendData(data []byte) error {
	if r.wsClient == nil {
		return errors.New("not connect")
	}
	return r.wsClient.SendData(data)
}

func (r *RobotDyj) HeartbeatInterval() time.Duration {
	return time.Second * 30
}

func (r *RobotDyj) HeartbeatData() []byte {
	req := &pbgo.Ping{Timestamp: time.Now().Unix()}
	data, err := r.msgProcessor.Marshal(req)
	if err != nil {
		return nil
	}
	return data
}

func (r *RobotDyj) HandleMessage(data []byte) {
	msg, err := r.msgProcessor.Unmarshal(data)
	if err != nil {
		log.Error("unmarshal msg err: ", err)
		return
	}

	if r.callback != nil {
		r.callback.OnMessage(msg)
	}
}
