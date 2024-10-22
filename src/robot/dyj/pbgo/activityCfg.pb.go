// protoc --proto_path=./pb  --proto_path=./pb/ --go_out=./pbgo --go_opt=paths=source_relative ./pb/activityCfg.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: activityCfg.proto

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

// 生效类型
type EActEffectType int32

const (
	EActEffectType_EAET_Always      EActEffectType = 0 //常驻
	EActEffectType_EAET_TimeSection EActEffectType = 1 //时间段
	EActEffectType_EAET_Period      EActEffectType = 2 //周期
	EActEffectType_EAET_NewUser     EActEffectType = 3 //新用户
)

// Enum value maps for EActEffectType.
var (
	EActEffectType_name = map[int32]string{
		0: "EAET_Always",
		1: "EAET_TimeSection",
		2: "EAET_Period",
		3: "EAET_NewUser",
	}
	EActEffectType_value = map[string]int32{
		"EAET_Always":      0,
		"EAET_TimeSection": 1,
		"EAET_Period":      2,
		"EAET_NewUser":     3,
	}
)

func (x EActEffectType) Enum() *EActEffectType {
	p := new(EActEffectType)
	*p = x
	return p
}

func (x EActEffectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EActEffectType) Descriptor() protoreflect.EnumDescriptor {
	return file_activityCfg_proto_enumTypes[0].Descriptor()
}

func (EActEffectType) Type() protoreflect.EnumType {
	return &file_activityCfg_proto_enumTypes[0]
}

func (x EActEffectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EActEffectType.Descriptor instead.
func (EActEffectType) EnumDescriptor() ([]byte, []int) {
	return file_activityCfg_proto_rawDescGZIP(), []int{0}
}

type ActTimeSection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	No            uint32 `protobuf:"varint,1,opt,name=No,proto3" json:"No,omitempty"`
	BeginTime     uint32 `protobuf:"varint,2,opt,name=BeginTime,proto3" json:"BeginTime,omitempty"`
	EndTime       uint32 `protobuf:"varint,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	EndRewardTime uint32 `protobuf:"varint,4,opt,name=EndRewardTime,proto3" json:"EndRewardTime,omitempty"`
}

func (x *ActTimeSection) Reset() {
	*x = ActTimeSection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activityCfg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActTimeSection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActTimeSection) ProtoMessage() {}

func (x *ActTimeSection) ProtoReflect() protoreflect.Message {
	mi := &file_activityCfg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActTimeSection.ProtoReflect.Descriptor instead.
func (*ActTimeSection) Descriptor() ([]byte, []int) {
	return file_activityCfg_proto_rawDescGZIP(), []int{0}
}

func (x *ActTimeSection) GetNo() uint32 {
	if x != nil {
		return x.No
	}
	return 0
}

func (x *ActTimeSection) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *ActTimeSection) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ActTimeSection) GetEndRewardTime() uint32 {
	if x != nil {
		return x.EndRewardTime
	}
	return 0
}

type ActivityCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"` // id
	Type uint32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	// string Name = 3;    // 不读了, csv不是utf8会报错
	OpenLvl       uint32         `protobuf:"varint,3,opt,name=OpenLvl,proto3" json:"OpenLvl,omitempty"`
	BeginTime     uint32         `protobuf:"varint,4,opt,name=BeginTime,proto3" json:"BeginTime,omitempty"`
	EndTime       uint32         `protobuf:"varint,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	EndRewardTime uint32         `protobuf:"varint,6,opt,name=EndRewardTime,proto3" json:"EndRewardTime,omitempty"`
	RankType      uint32         `protobuf:"varint,10,opt,name=RankType,proto3" json:"RankType,omitempty"`
	TimeNo        uint32         `protobuf:"varint,11,opt,name=TimeNo,proto3" json:"TimeNo,omitempty"`                                   //时间编号
	EffectType    EActEffectType `protobuf:"varint,12,opt,name=EffectType,proto3,enum=model.EActEffectType" json:"EffectType,omitempty"` //生效类型
	FlushId       EFlushId       `protobuf:"varint,13,opt,name=FlushId,proto3,enum=model.EFlushId" json:"FlushId,omitempty"`             //刷新id
	Duration      int64          `protobuf:"varint,14,opt,name=Duration,proto3" json:"Duration,omitempty"`                               //持续时间
	RewardEmailId uint32         `protobuf:"varint,15,opt,name=RewardEmailId,proto3" json:"RewardEmailId,omitempty"`                     //奖励邮件id
}

func (x *ActivityCfg) Reset() {
	*x = ActivityCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activityCfg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityCfg) ProtoMessage() {}

func (x *ActivityCfg) ProtoReflect() protoreflect.Message {
	mi := &file_activityCfg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityCfg.ProtoReflect.Descriptor instead.
func (*ActivityCfg) Descriptor() ([]byte, []int) {
	return file_activityCfg_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityCfg) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ActivityCfg) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ActivityCfg) GetOpenLvl() uint32 {
	if x != nil {
		return x.OpenLvl
	}
	return 0
}

func (x *ActivityCfg) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *ActivityCfg) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ActivityCfg) GetEndRewardTime() uint32 {
	if x != nil {
		return x.EndRewardTime
	}
	return 0
}

func (x *ActivityCfg) GetRankType() uint32 {
	if x != nil {
		return x.RankType
	}
	return 0
}

func (x *ActivityCfg) GetTimeNo() uint32 {
	if x != nil {
		return x.TimeNo
	}
	return 0
}

func (x *ActivityCfg) GetEffectType() EActEffectType {
	if x != nil {
		return x.EffectType
	}
	return EActEffectType_EAET_Always
}

func (x *ActivityCfg) GetFlushId() EFlushId {
	if x != nil {
		return x.FlushId
	}
	return EFlushId_EFI_None
}

func (x *ActivityCfg) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *ActivityCfg) GetRewardEmailId() uint32 {
	if x != nil {
		return x.RewardEmailId
	}
	return 0
}

var File_activityCfg_proto protoreflect.FileDescriptor

var file_activityCfg_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x66, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x4e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x4e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x45, 0x6e, 0x64, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x81, 0x03, 0x0a, 0x0b, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x66, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x70, 0x65, 0x6e, 0x4c, 0x76, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4f,
	0x70, 0x65, 0x6e, 0x4c, 0x76, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x12, 0x35, 0x0a, 0x0a, 0x45, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x41, 0x63, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x29, 0x0a, 0x07, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x49,
	0x64, 0x52, 0x07, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x2a, 0x5a, 0x0a, 0x0e,
	0x45, 0x41, 0x63, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x45, 0x41, 0x45, 0x54, 0x5f, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x41, 0x45, 0x54, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x41, 0x45, 0x54, 0x5f, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x41, 0x45, 0x54, 0x5f, 0x4e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x10, 0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activityCfg_proto_rawDescOnce sync.Once
	file_activityCfg_proto_rawDescData = file_activityCfg_proto_rawDesc
)

func file_activityCfg_proto_rawDescGZIP() []byte {
	file_activityCfg_proto_rawDescOnce.Do(func() {
		file_activityCfg_proto_rawDescData = protoimpl.X.CompressGZIP(file_activityCfg_proto_rawDescData)
	})
	return file_activityCfg_proto_rawDescData
}

var file_activityCfg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_activityCfg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_activityCfg_proto_goTypes = []interface{}{
	(EActEffectType)(0),    // 0: model.EActEffectType
	(*ActTimeSection)(nil), // 1: model.ActTimeSection
	(*ActivityCfg)(nil),    // 2: model.ActivityCfg
	(EFlushId)(0),          // 3: model.EFlushId
}
var file_activityCfg_proto_depIdxs = []int32{
	0, // 0: model.ActivityCfg.EffectType:type_name -> model.EActEffectType
	3, // 1: model.ActivityCfg.FlushId:type_name -> model.EFlushId
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_activityCfg_proto_init() }
func file_activityCfg_proto_init() {
	if File_activityCfg_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_activityCfg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActTimeSection); i {
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
		file_activityCfg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityCfg); i {
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
			RawDescriptor: file_activityCfg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_activityCfg_proto_goTypes,
		DependencyIndexes: file_activityCfg_proto_depIdxs,
		EnumInfos:         file_activityCfg_proto_enumTypes,
		MessageInfos:      file_activityCfg_proto_msgTypes,
	}.Build()
	File_activityCfg_proto = out.File
	file_activityCfg_proto_rawDesc = nil
	file_activityCfg_proto_goTypes = nil
	file_activityCfg_proto_depIdxs = nil
}
