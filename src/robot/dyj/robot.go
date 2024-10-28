package dyj

import (
	"RobotTool/src/app/api"
	"RobotTool/src/common"
	"RobotTool/src/jsdef/jsmsg"
	"RobotTool/src/network"
	"RobotTool/src/robot/dyj/pbgo"
	"errors"
	"github.com/magicsea/behavior3go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type RobotDyj struct {
	wsClient     *network.WsClient
	msgProcessor *MsgProcessor
	appApi       api.IAppApi

	msgCounter atomic.Uint64
	lock       sync.Mutex
}

func NewRobotDyj() *RobotDyj {
	rob := &RobotDyj{
		msgProcessor: NewMsgProcessor(),
	}

	registerWsMsg(rob.msgProcessor)

	return rob
}

// ******对app接口******//
func (r *RobotDyj) RegisterBevMap(mp *behavior3go.RegisterStructMaps) {
	mp.Register("Login", &Login{})
}

func (r *RobotDyj) SetAppApi(appApi api.IAppApi) {
	r.appApi = appApi
}

func (r *RobotDyj) GetProtoRequest() []*jsmsg.ProtoMsg {
	ret := make([]*jsmsg.ProtoMsg, 0)

	allMsg := r.msgProcessor.GetAllMsgInfo()
	for id, tp := range allMsg {
		rvTp := tp.ProtoReflect()
		name := string(rvTp.Descriptor().Name())
		if !strings.Contains(name, "C2S") {
			continue
		}

		schema, err := os.ReadFile("./src/robot/dyj/schema/" + name + ".json")
		if err != nil {
			log.Warning(err)
			continue
		}

		defalutVal := rvTp.New().Interface()
		err = common.InitPbMessage(defalutVal)
		if err != nil {
			log.Warning(err)
			continue
		}

		msg := &jsmsg.ProtoMsg{
			Id:           id,
			Name:         name,
			Schema:       common.NewJsonRawWrapper(schema),
			DefaultValue: common.NewJsonProtoWrapper(defalutVal),
		}
		ret = append(ret, msg)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Id.(uint16) < ret[j].Id.(uint16)
	})

	return ret
}

//******对app接口******//

func (r *RobotDyj) Connect(url string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.wsClient != nil {
		return errors.New("already connect")
	}

	r.msgCounter.Store(0)

	client := network.NewWsClient(r)
	err := client.Connect(url)
	if err != nil {
		return err
	}

	r.wsClient = client

	return nil
}

func (r *RobotDyj) Close() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.wsClient == nil {
		return nil
	}

	_ = r.wsClient.Close()
	r.wsClient = nil
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

	err = r.SendData(data)
	if err != nil {
		r.addMsgToUi(msg, "SEND FAILED")
		return err
	}

	r.addMsgToUi(msg, "OK")
	return nil
}

func (r *RobotDyj) SendData(data []byte) error {
	r.lock.Lock()
	if r.wsClient == nil {
		r.lock.Unlock()
		return errors.New("not connect")
	}
	r.lock.Unlock()
	return r.wsClient.SendData(data)
}

func (r *RobotDyj) HeartbeatInterval() time.Duration {
	return time.Second * 30
}

func (r *RobotDyj) OnHeartbeat() {
	req := &pbgo.Ping{Timestamp: time.Now().Unix()}
	_ = r.SendMsg(req)
}

func (r *RobotDyj) HandleMessage(data []byte) {
	rsp, err := r.msgProcessor.Unmarshal(data)
	if err != nil {
		log.Error("unmarshal msg err: ", err)
		return
	}

	r.addMsgToUi(rsp, "OK")
}

func (r *RobotDyj) OnClose() {
	log.Info("ws client close")
}

func (r *RobotDyj) addMsgToUi(msg proto.Message, ret string) {
	if r.appApi == nil {
		return
	}

	rtRsp := reflect.TypeOf(msg)
	showMsg := &jsmsg.ShowMessage{
		Id:        r.msgCounter.Add(1),
		ProtoId:   r.msgProcessor.GetMsgId(msg),
		ProtoType: rtRsp.Elem().Name(),
		Msg:       msg,
		Type:      jsmsg.Success,
		Result:    ret,
	}

	rvRsp := reflect.ValueOf(msg).Elem()
	if val := rvRsp.FieldByName("ErrCode"); val.IsValid() {
		showMsg.Result = pbgo.STATUS(val.Int()).String()
		if val.Int() != 0 {
			showMsg.Type = jsmsg.Danger
		}
		r.appApi.AddTipMsg(showMsg.Result)
	}

	if logonRsp, ok := msg.(*pbgo.S2CLogin); ok {
		r.appApi.AddTipTag(&jsmsg.TipTag{
			Name: "uin",
			Type: jsmsg.Success,
			Msg:  logonRsp.UserBase.Uid,
		})
		r.appApi.AddTipTag(&jsmsg.TipTag{
			Name: "idNum",
			Type: jsmsg.Success,
			Msg:  strconv.Itoa(int(logonRsp.UserBase.IdNum)),
		})
		r.appApi.AddTipTag(&jsmsg.TipTag{
			Name: "nick",
			Type: jsmsg.Success,
			Msg:  logonRsp.UserBase.NickName,
		})
	}

	r.appApi.AddShowMsg(showMsg)
}
