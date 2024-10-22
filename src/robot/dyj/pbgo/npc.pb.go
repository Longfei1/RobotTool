// protoc --proto_path=./pb --go_out=./pbgo --go_opt=paths=source_relative ./pb/npc.proto
// protoc-go-inject-tag -input="./pbgo/npc.pb.go"
// goctl model mongo -e -dir ./model -t Npc --home template -c
//!!!!!!! 删除生成的worldtypes, 代码成替换成model对象

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: npc.proto

package pbgo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NpcOperaId int32

const (
	NpcOperaId_NOI_Invalid NpcOperaId = 0
	NpcOperaId_NOI_Invite  NpcOperaId = 1
	NpcOperaId_NOI_Feed    NpcOperaId = 2
)

// Enum value maps for NpcOperaId.
var (
	NpcOperaId_name = map[int32]string{
		0: "NOI_Invalid",
		1: "NOI_Invite",
		2: "NOI_Feed",
	}
	NpcOperaId_value = map[string]int32{
		"NOI_Invalid": 0,
		"NOI_Invite":  1,
		"NOI_Feed":    2,
	}
)

func (x NpcOperaId) Enum() *NpcOperaId {
	p := new(NpcOperaId)
	*p = x
	return p
}

func (x NpcOperaId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NpcOperaId) Descriptor() protoreflect.EnumDescriptor {
	return file_npc_proto_enumTypes[0].Descriptor()
}

func (NpcOperaId) Type() protoreflect.EnumType {
	return &file_npc_proto_enumTypes[0]
}

func (x NpcOperaId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NpcOperaId.Descriptor instead.
func (NpcOperaId) EnumDescriptor() ([]byte, []int) {
	return file_npc_proto_rawDescGZIP(), []int{0}
}

type NPCInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"npcId"
	NpcId uint32 `protobuf:"varint,1,opt,name=npcId,proto3" json:"npcId,omitempty"`
	// @inject_tag: bson:"favorAbility"
	FavorAbility uint32 `protobuf:"varint,2,opt,name=favorAbility,proto3" json:"favorAbility,omitempty"` //好感度
	// @inject_tag: bson:"npcDayInfo"
	NpcDayInfo *NPCDayInfo `protobuf:"bytes,3,opt,name=npcDayInfo,proto3" json:"npcDayInfo,omitempty"`
	// @inject_tag: bson:"rewardedTokenList"
	RewardedTokenList []uint32 `protobuf:"varint,4,rep,packed,name=rewardedTokenList,proto3" json:"rewardedTokenList,omitempty"` // 已领取信物奖励的信物id列表
	// @inject_tag: bson:"firstTalk"
	FirstTalk bool `protobuf:"varint,5,opt,name=firstTalk,proto3" json:"firstTalk,omitempty"` //是否进行初见任务
}

func (x *NPCInfo) Reset() {
	*x = NPCInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NPCInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NPCInfo) ProtoMessage() {}

func (x *NPCInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NPCInfo.ProtoReflect.Descriptor instead.
func (*NPCInfo) Descriptor() ([]byte, []int) {
	return file_npc_proto_rawDescGZIP(), []int{0}
}

func (x *NPCInfo) GetNpcId() uint32 {
	if x != nil {
		return x.NpcId
	}
	return 0
}

func (x *NPCInfo) GetFavorAbility() uint32 {
	if x != nil {
		return x.FavorAbility
	}
	return 0
}

func (x *NPCInfo) GetNpcDayInfo() *NPCDayInfo {
	if x != nil {
		return x.NpcDayInfo
	}
	return nil
}

func (x *NPCInfo) GetRewardedTokenList() []uint32 {
	if x != nil {
		return x.RewardedTokenList
	}
	return nil
}

func (x *NPCInfo) GetFirstTalk() bool {
	if x != nil {
		return x.FirstTalk
	}
	return false
}

// npc每天的操作记录
type NPCOperaRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"time"
	Time uint32 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"` // 操作时间
	// @inject_tag: bson:"operaId"
	OperaId NpcOperaId `protobuf:"varint,2,opt,name=operaId,proto3,enum=model.NpcOperaId" json:"operaId,omitempty"` // 操作id
	// @inject_tag: bson:"operaValue"
	OperaValue uint32 `protobuf:"varint,3,opt,name=operaValue,proto3" json:"operaValue,omitempty"` // 操作对应的值，比如增加好感度等
}

func (x *NPCOperaRecord) Reset() {
	*x = NPCOperaRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NPCOperaRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NPCOperaRecord) ProtoMessage() {}

func (x *NPCOperaRecord) ProtoReflect() protoreflect.Message {
	mi := &file_npc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NPCOperaRecord.ProtoReflect.Descriptor instead.
func (*NPCOperaRecord) Descriptor() ([]byte, []int) {
	return file_npc_proto_rawDescGZIP(), []int{1}
}

func (x *NPCOperaRecord) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *NPCOperaRecord) GetOperaId() NpcOperaId {
	if x != nil {
		return x.OperaId
	}
	return NpcOperaId_NOI_Invalid
}

func (x *NPCOperaRecord) GetOperaValue() uint32 {
	if x != nil {
		return x.OperaValue
	}
	return 0
}

// npc每天需要刷新的信息
type NPCDayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"lastInvitedTime"
	LastInvitedTime uint32 `protobuf:"varint,1,opt,name=lastInvitedTime,proto3" json:"lastInvitedTime,omitempty"` //上次被邀请时间
	// @inject_tag: bson:"feedTime"
	FeedTime uint32 `protobuf:"varint,2,opt,name=feedTime,proto3" json:"feedTime,omitempty"` // 被投喂次数
	// @inject_tag: bson:"talkFlag"
	TalkFlag bool `protobuf:"varint,3,opt,name=talkFlag,proto3" json:"talkFlag,omitempty"` // 今日是否谈天说地的标志位
	// @inject_tag: bson:"operaRecordList"
	OperaRecordList []*NPCOperaRecord `protobuf:"bytes,4,rep,name=operaRecordList,proto3" json:"operaRecordList,omitempty"`
}

func (x *NPCDayInfo) Reset() {
	*x = NPCDayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NPCDayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NPCDayInfo) ProtoMessage() {}

func (x *NPCDayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NPCDayInfo.ProtoReflect.Descriptor instead.
func (*NPCDayInfo) Descriptor() ([]byte, []int) {
	return file_npc_proto_rawDescGZIP(), []int{2}
}

func (x *NPCDayInfo) GetLastInvitedTime() uint32 {
	if x != nil {
		return x.LastInvitedTime
	}
	return 0
}

func (x *NPCDayInfo) GetFeedTime() uint32 {
	if x != nil {
		return x.FeedTime
	}
	return 0
}

func (x *NPCDayInfo) GetTalkFlag() bool {
	if x != nil {
		return x.TalkFlag
	}
	return false
}

func (x *NPCDayInfo) GetOperaRecordList() []*NPCOperaRecord {
	if x != nil {
		return x.OperaRecordList
	}
	return nil
}

type Npc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"_id"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// @inject_tag: bson:"npcList"
	NpcList []*NPCInfo `protobuf:"bytes,2,rep,name=npcList,proto3" json:"npcList,omitempty"`
	// @inject_tag: bson:"seq"
	Seq uint64 `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Npc) Reset() {
	*x = Npc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Npc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Npc) ProtoMessage() {}

func (x *Npc) ProtoReflect() protoreflect.Message {
	mi := &file_npc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Npc.ProtoReflect.Descriptor instead.
func (*Npc) Descriptor() ([]byte, []int) {
	return file_npc_proto_rawDescGZIP(), []int{3}
}

func (x *Npc) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Npc) GetNpcList() []*NPCInfo {
	if x != nil {
		return x.NpcList
	}
	return nil
}

func (x *Npc) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

var File_npc_proto protoreflect.FileDescriptor

var file_npc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0xc2, 0x01, 0x0a, 0x07, 0x4e, 0x50, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x70, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e,
	0x70, 0x63, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x0a, 0x6e, 0x70, 0x63, 0x44,
	0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x50, 0x43, 0x44, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x6e, 0x70, 0x63, 0x44, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x54, 0x61, 0x6c, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x54, 0x61, 0x6c, 0x6b, 0x22, 0x71, 0x0a, 0x0e, 0x4e, 0x50, 0x43, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a,
	0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x70, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x49,
	0x64, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0a, 0x4e,
	0x50, 0x43, 0x44, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x61, 0x6c, 0x6b, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x74, 0x61, 0x6c, 0x6b, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x3f, 0x0a, 0x0f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x50, 0x43,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x03,
	0x4e, 0x70, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x6e, 0x70, 0x63, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e,
	0x50, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6e, 0x70, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65,
	0x71, 0x2a, 0x3b, 0x0a, 0x0a, 0x4e, 0x70, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12,
	0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x49, 0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x49, 0x5f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x49, 0x5f, 0x46, 0x65, 0x65, 0x64, 0x10, 0x02, 0x42, 0x0d,
	0x5a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npc_proto_rawDescOnce sync.Once
	file_npc_proto_rawDescData = file_npc_proto_rawDesc
)

func file_npc_proto_rawDescGZIP() []byte {
	file_npc_proto_rawDescOnce.Do(func() {
		file_npc_proto_rawDescData = protoimpl.X.CompressGZIP(file_npc_proto_rawDescData)
	})
	return file_npc_proto_rawDescData
}

var file_npc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npc_proto_goTypes = []interface{}{
	(NpcOperaId)(0),        // 0: model.NpcOperaId
	(*NPCInfo)(nil),        // 1: model.NPCInfo
	(*NPCOperaRecord)(nil), // 2: model.NPCOperaRecord
	(*NPCDayInfo)(nil),     // 3: model.NPCDayInfo
	(*Npc)(nil),            // 4: model.Npc
}
var file_npc_proto_depIdxs = []int32{
	3, // 0: model.NPCInfo.npcDayInfo:type_name -> model.NPCDayInfo
	0, // 1: model.NPCOperaRecord.operaId:type_name -> model.NpcOperaId
	2, // 2: model.NPCDayInfo.operaRecordList:type_name -> model.NPCOperaRecord
	1, // 3: model.Npc.npcList:type_name -> model.NPCInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npc_proto_init() }
func file_npc_proto_init() {
	if File_npc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NPCInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NPCOperaRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NPCDayInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Npc); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npc_proto_goTypes,
		DependencyIndexes: file_npc_proto_depIdxs,
		EnumInfos:         file_npc_proto_enumTypes,
		MessageInfos:      file_npc_proto_msgTypes,
	}.Build()
	File_npc_proto = out.File
	file_npc_proto_rawDesc = nil
	file_npc_proto_goTypes = nil
	file_npc_proto_depIdxs = nil
}
