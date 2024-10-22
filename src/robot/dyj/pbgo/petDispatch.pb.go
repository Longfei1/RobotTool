// protoc --proto_path=./pb --go_out=./pbgo --go_opt=paths=source_relative ./pb/petDispatch.proto
// protoc-go-inject-tag -input="./pbgo/petDispatch.pb.go"
// goctl model mongo -e -dir ./model -t PetDispatch --home template -c

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: petDispatch.proto

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

type PetDispatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"pointId"
	PointId uint32 `protobuf:"varint,1,opt,name=pointId,proto3" json:"pointId,omitempty"` // 派遣点id
	// @inject_tag: bson:"taskId"
	TaskId uint32 `protobuf:"varint,2,opt,name=taskId,proto3" json:"taskId,omitempty"` // 派遣任务id 这个任务是单独的
	// @inject_tag: bson:"dispatchEndTime"
	DispatchEndTime uint32 `protobuf:"varint,3,opt,name=dispatchEndTime,proto3" json:"dispatchEndTime,omitempty"`
	// @inject_tag: bson:"petUidList"
	PetUidList []string `protobuf:"bytes,4,rep,name=petUidList,proto3" json:"petUidList,omitempty"` // 是否派遣宠物
	// @inject_tag: bson:"rewardFlag"
	RewardFlag bool `protobuf:"varint,5,opt,name=rewardFlag,proto3" json:"rewardFlag,omitempty"` // 是否领取奖励
}

func (x *PetDispatchInfo) Reset() {
	*x = PetDispatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petDispatch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetDispatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetDispatchInfo) ProtoMessage() {}

func (x *PetDispatchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_petDispatch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetDispatchInfo.ProtoReflect.Descriptor instead.
func (*PetDispatchInfo) Descriptor() ([]byte, []int) {
	return file_petDispatch_proto_rawDescGZIP(), []int{0}
}

func (x *PetDispatchInfo) GetPointId() uint32 {
	if x != nil {
		return x.PointId
	}
	return 0
}

func (x *PetDispatchInfo) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *PetDispatchInfo) GetDispatchEndTime() uint32 {
	if x != nil {
		return x.DispatchEndTime
	}
	return 0
}

func (x *PetDispatchInfo) GetPetUidList() []string {
	if x != nil {
		return x.PetUidList
	}
	return nil
}

func (x *PetDispatchInfo) GetRewardFlag() bool {
	if x != nil {
		return x.RewardFlag
	}
	return false
}

type PetDispatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"_id"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// @inject_tag: bson:"lastRefreshTime"
	LastRefreshTime uint32 `protobuf:"varint,2,opt,name=lastRefreshTime,proto3" json:"lastRefreshTime,omitempty"`
	// @inject_tag: bson:"petDispatchInfos"
	PetDispatchInfos []*PetDispatchInfo `protobuf:"bytes,3,rep,name=petDispatchInfos,proto3" json:"petDispatchInfos,omitempty"`
	// @inject_tag: bson:"seq"
	Seq uint64 `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *PetDispatch) Reset() {
	*x = PetDispatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petDispatch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetDispatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetDispatch) ProtoMessage() {}

func (x *PetDispatch) ProtoReflect() protoreflect.Message {
	mi := &file_petDispatch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetDispatch.ProtoReflect.Descriptor instead.
func (*PetDispatch) Descriptor() ([]byte, []int) {
	return file_petDispatch_proto_rawDescGZIP(), []int{1}
}

func (x *PetDispatch) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PetDispatch) GetLastRefreshTime() uint32 {
	if x != nil {
		return x.LastRefreshTime
	}
	return 0
}

func (x *PetDispatch) GetPetDispatchInfos() []*PetDispatchInfo {
	if x != nil {
		return x.PetDispatchInfos
	}
	return nil
}

func (x *PetDispatch) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

var File_petDispatch_proto protoreflect.FileDescriptor

var file_petDispatch_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xad, 0x01, 0x0a, 0x0f, 0x50,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x74, 0x55, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x65, 0x74, 0x55, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x50,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x70, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x70, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65,
	0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71, 0x42, 0x0d, 0x5a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_petDispatch_proto_rawDescOnce sync.Once
	file_petDispatch_proto_rawDescData = file_petDispatch_proto_rawDesc
)

func file_petDispatch_proto_rawDescGZIP() []byte {
	file_petDispatch_proto_rawDescOnce.Do(func() {
		file_petDispatch_proto_rawDescData = protoimpl.X.CompressGZIP(file_petDispatch_proto_rawDescData)
	})
	return file_petDispatch_proto_rawDescData
}

var file_petDispatch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_petDispatch_proto_goTypes = []interface{}{
	(*PetDispatchInfo)(nil), // 0: model.PetDispatchInfo
	(*PetDispatch)(nil),     // 1: model.PetDispatch
}
var file_petDispatch_proto_depIdxs = []int32{
	0, // 0: model.PetDispatch.petDispatchInfos:type_name -> model.PetDispatchInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_petDispatch_proto_init() }
func file_petDispatch_proto_init() {
	if File_petDispatch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petDispatch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetDispatchInfo); i {
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
		file_petDispatch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetDispatch); i {
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
			RawDescriptor: file_petDispatch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_petDispatch_proto_goTypes,
		DependencyIndexes: file_petDispatch_proto_depIdxs,
		MessageInfos:      file_petDispatch_proto_msgTypes,
	}.Build()
	File_petDispatch_proto = out.File
	file_petDispatch_proto_rawDesc = nil
	file_petDispatch_proto_goTypes = nil
	file_petDispatch_proto_depIdxs = nil
}