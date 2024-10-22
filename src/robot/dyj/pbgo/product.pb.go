// protoc --proto_path=./pb --go_out=./pbgo --go_opt=paths=source_relative ./pb/product.proto
// protoc-go-inject-tag -input="./pbgo/product.pb.go"
// goctl model mongo -e -dir ./model -t Product --home template -c
//!!!!!!! 删除生成的worldtypes, 代码成替换成model对象

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: product.proto

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

type ProductItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"globalItemId"
	GlobalItemId uint32 `protobuf:"varint,1,opt,name=globalItemId,proto3" json:"globalItemId,omitempty"`
	// @inject_tag: bson:"startTime"
	StartTime uint32 `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"` //! 时间戳(秒)
	// @inject_tag: bson:"endTime"
	EndTime uint32 `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"` //! 时间戳(秒)
}

func (x *ProductItem) Reset() {
	*x = ProductItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductItem) ProtoMessage() {}

func (x *ProductItem) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductItem.ProtoReflect.Descriptor instead.
func (*ProductItem) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductItem) GetGlobalItemId() uint32 {
	if x != nil {
		return x.GlobalItemId
	}
	return 0
}

func (x *ProductItem) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *ProductItem) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type ProductInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"buildUid"
	BuildUid string `protobuf:"bytes,1,opt,name=buildUid,proto3" json:"buildUid,omitempty"`
	// @inject_tag: bson:"productList"
	ProductList []*ProductItem `protobuf:"bytes,2,rep,name=productList,proto3" json:"productList,omitempty"` //! 长度一定与格子数量相等
	// @inject_tag: bson:"slotNum"
	SlotNum uint32 `protobuf:"varint,3,opt,name=slotNum,proto3" json:"slotNum,omitempty"`
}

func (x *ProductInfo) Reset() {
	*x = ProductInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInfo) ProtoMessage() {}

func (x *ProductInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInfo.ProtoReflect.Descriptor instead.
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductInfo) GetBuildUid() string {
	if x != nil {
		return x.BuildUid
	}
	return ""
}

func (x *ProductInfo) GetProductList() []*ProductItem {
	if x != nil {
		return x.ProductList
	}
	return nil
}

func (x *ProductInfo) GetSlotNum() uint32 {
	if x != nil {
		return x.SlotNum
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"_id"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// @inject_tag: bson:"productList"
	ProductInfoList []*ProductInfo `protobuf:"bytes,2,rep,name=productInfoList,proto3" json:"productInfoList,omitempty"`
	// @inject_tag: bson:"seq"
	Seq uint64 `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Product) GetProductInfoList() []*ProductInfo {
	if x != nil {
		return x.ProductInfoList
	}
	return nil
}

func (x *Product) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x69, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x79, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x55, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x6b, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x0f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71, 0x42, 0x0d, 0x5a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_product_proto_goTypes = []interface{}{
	(*ProductItem)(nil), // 0: model.ProductItem
	(*ProductInfo)(nil), // 1: model.ProductInfo
	(*Product)(nil),     // 2: model.Product
}
var file_product_proto_depIdxs = []int32{
	0, // 0: model.ProductInfo.productList:type_name -> model.ProductItem
	1, // 1: model.Product.productInfoList:type_name -> model.ProductInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductItem); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInfo); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}