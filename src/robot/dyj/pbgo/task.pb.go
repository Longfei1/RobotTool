// protoc --proto_path=./pb --go_out=./pbgo --go_opt=paths=source_relative ./pb/task.proto
// protoc-go-inject-tag -input="./pbgo/task.pb.go"
// goctl model mongo -e -dir ./model -t Task --home template -c
//!!!!!!! 删除生成的worldtypes, 代码成替换成model对象

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: task.proto

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

// 用于redis的消息队列的数据载体
type TaskCondDataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NeedEventId  ETaskEventId  `protobuf:"varint,1,opt,name=needEventId,proto3,enum=model.ETaskEventId" json:"needEventId,omitempty"`
	TaskCondData *TaskCondData `protobuf:"bytes,2,opt,name=TaskCondData,proto3" json:"TaskCondData,omitempty"`
}

func (x *TaskCondDataInfo) Reset() {
	*x = TaskCondDataInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCondDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCondDataInfo) ProtoMessage() {}

func (x *TaskCondDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCondDataInfo.ProtoReflect.Descriptor instead.
func (*TaskCondDataInfo) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskCondDataInfo) GetNeedEventId() ETaskEventId {
	if x != nil {
		return x.NeedEventId
	}
	return ETaskEventId_ETEI_Invalid
}

func (x *TaskCondDataInfo) GetTaskCondData() *TaskCondData {
	if x != nil {
		return x.TaskCondData
	}
	return nil
}

type TaskCondData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"globalItemId"
	GlobalItemId uint32 `protobuf:"varint,2,opt,name=globalItemId,proto3" json:"globalItemId,omitempty"`
	// @inject_tag: bson:"num"
	Num uint32 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *TaskCondData) Reset() {
	*x = TaskCondData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCondData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCondData) ProtoMessage() {}

func (x *TaskCondData) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCondData.ProtoReflect.Descriptor instead.
func (*TaskCondData) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskCondData) GetGlobalItemId() uint32 {
	if x != nil {
		return x.GlobalItemId
	}
	return 0
}

func (x *TaskCondData) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// index为uid和taskId，唯一id系统生成
	// @inject_tag: bson:"uid"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// @inject_tag: bson:"taskId"
	TaskId uint32 `protobuf:"varint,2,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// @inject_tag: bson:"updateAt"
	UpdateAt uint32 `protobuf:"varint,3,opt,name=updateAt,proto3" json:"updateAt,omitempty"` //! 更新时间
	// @inject_tag: bson:"taskDataList"
	TaskDataList []*TaskCondData `protobuf:"bytes,4,rep,name=taskDataList,proto3" json:"taskDataList,omitempty"` //! 完成条件的数据
	// @inject_tag: bson:"taskStatus"
	TaskStatus ETaskStatus `protobuf:"varint,5,opt,name=taskStatus,proto3,enum=model.ETaskStatus" json:"taskStatus,omitempty"`
	// @inject_tag: bson:"needEventId"
	NeedEventId ETaskEventId `protobuf:"varint,6,opt,name=needEventId,proto3,enum=model.ETaskEventId" json:"needEventId,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Task) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Task) GetUpdateAt() uint32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Task) GetTaskDataList() []*TaskCondData {
	if x != nil {
		return x.TaskDataList
	}
	return nil
}

func (x *Task) GetTaskStatus() ETaskStatus {
	if x != nil {
		return x.TaskStatus
	}
	return ETaskStatus_ETS_Invalid
}

func (x *Task) GetNeedEventId() ETaskEventId {
	if x != nil {
		return x.NeedEventId
	}
	return ETaskEventId_ETEI_Invalid
}

type TaskChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid            string          `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TaskChangeList []*TaskCondData `protobuf:"bytes,2,rep,name=taskChangeList,proto3" json:"taskChangeList,omitempty"` //! 条件的数据
}

func (x *TaskChange) Reset() {
	*x = TaskChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskChange) ProtoMessage() {}

func (x *TaskChange) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskChange.ProtoReflect.Descriptor instead.
func (*TaskChange) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskChange) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TaskChange) GetTaskChangeList() []*TaskCondData {
	if x != nil {
		return x.TaskChangeList
	}
	return nil
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a,
	0x0c, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x43, 0x6f, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0xf0, 0x01, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x74,
	0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x74, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x6e, 0x65, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x5b, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x3b, 0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x74, 0x61,
	0x73, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0d, 0x5a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_task_proto_goTypes = []interface{}{
	(*TaskCondDataInfo)(nil), // 0: model.TaskCondDataInfo
	(*TaskCondData)(nil),     // 1: model.TaskCondData
	(*Task)(nil),             // 2: model.Task
	(*TaskChange)(nil),       // 3: model.TaskChange
	(ETaskEventId)(0),        // 4: model.ETaskEventId
	(ETaskStatus)(0),         // 5: model.ETaskStatus
}
var file_task_proto_depIdxs = []int32{
	4, // 0: model.TaskCondDataInfo.needEventId:type_name -> model.ETaskEventId
	1, // 1: model.TaskCondDataInfo.TaskCondData:type_name -> model.TaskCondData
	1, // 2: model.Task.taskDataList:type_name -> model.TaskCondData
	5, // 3: model.Task.taskStatus:type_name -> model.ETaskStatus
	4, // 4: model.Task.needEventId:type_name -> model.ETaskEventId
	1, // 5: model.TaskChange.taskChangeList:type_name -> model.TaskCondData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCondDataInfo); i {
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
		file_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCondData); i {
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
		file_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskChange); i {
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
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
