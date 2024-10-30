package dyj

import (
	"encoding/binary"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"math"
	"reflect"
)

type MsgProcessor struct {
	littleEndian bool
	msgInfo      map[uint16]proto.Message
	msgID        map[reflect.Type]uint16
}

func NewMsgProcessor() *MsgProcessor {
	p := new(MsgProcessor)
	p.littleEndian = false
	p.msgID = make(map[reflect.Type]uint16)
	p.msgInfo = make(map[uint16]proto.Message)
	return p
}

func (p *MsgProcessor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

func (p *MsgProcessor) Register(msg proto.Message, eventType uint16) {
	msgType := reflect.TypeOf(msg)
	if _, ok := p.msgID[msgType]; ok {
		log.Errorf("message %s is already registered", msgType)
		return
	}
	if len(p.msgInfo) >= math.MaxUint16 {
		log.Errorf("too many protobuf messages (max = %v)", math.MaxUint16)
		return
	}

	p.msgInfo[eventType] = msg
	p.msgID[msgType] = eventType
}

func (p *MsgProcessor) GetMsgType(id uint16) proto.Message {
	return p.msgInfo[id]
}

func (p *MsgProcessor) GetMsgId(tp proto.Message) uint16 {
	return p.msgID[reflect.TypeOf(tp)]
}

func (p *MsgProcessor) GetAllMsgInfo() map[uint16]proto.Message {
	return p.msgInfo
}

func (p *MsgProcessor) Unmarshal(data []byte) (proto.Message, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short")
	}

	// id
	var id uint16
	if p.littleEndian {
		id = binary.LittleEndian.Uint16(data)
	} else {
		id = binary.BigEndian.Uint16(data)
	}
	msgType, ok := p.msgInfo[id]
	if !ok {
		return nil, fmt.Errorf("message id %v not registered", id)
	}

	// msg
	msg := msgType.ProtoReflect().New().Interface()
	return msg, proto.Unmarshal(data[2:], msg)
}

func (p *MsgProcessor) Marshal(msg proto.Message) ([]byte, error) {
	msgType := reflect.TypeOf(msg)

	// id
	_id, ok := p.msgID[msgType]
	if !ok {
		err := fmt.Errorf("message %v not registered", msgType)
		return nil, err
	}

	id := make([]byte, 2)
	if p.littleEndian {
		binary.LittleEndian.PutUint16(id, _id)
	} else {
		binary.BigEndian.PutUint16(id, _id)
	}

	// data
	data, err := proto.Marshal(msg)

	// 创建一个新的字节切片来存储 id 和 data
	combined := make([]byte, 2+len(data))

	// 将 id 复制到 combined 的前两个字节
	copy(combined[:2], id)

	// 将 data 复制到 combined 的剩余部分
	copy(combined[2:], data)
	return combined, err
}
