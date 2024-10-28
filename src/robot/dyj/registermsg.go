package dyj

import (
	"RobotTool/src/robot/dyj/pbgo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"strings"
)

var allTypeMap = make(map[string]proto.Message)

func init() {

}

func registerWsMsg(processor *MsgProcessor) {
	registerPckMsg(processor)

	registerSpecialMsg(processor)
}

func registerPckMsg(processor *MsgProcessor) {
	msgIdMap := pbgo.PCK_name

	for msgIdItem := range msgIdMap {
		if msgIdItem < 1000 {
			continue
		}

		name := msgIdMap[msgIdItem]
		name = strings.Trim(name, "_ID") // 去掉末尾的_ID
		msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName("msg." + name))
		if err != nil {
			log.Warn(err)
			continue
		}
		processor.Register(msgType.New().Interface(), uint16(msgIdItem))
	}
}

func registerSpecialMsg(processor *MsgProcessor) {
	processor.Register(&pbgo.Ping{}, 10000)
	processor.Register(&pbgo.Pong{}, 10001)
}
