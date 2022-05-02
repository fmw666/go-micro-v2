// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: orderModel.proto

package service

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

type OrderModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// @inject_tag: json:"user_id"
	UserID uint32 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	// @inject_tag: json:"created_at"
	CreatedAt int64 `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	// @inject_tag: json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	// @inject_tag: json:"deleted_at"
	DeletedAt int64 `protobuf:"varint,5,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *OrderModel) Reset() {
	*x = OrderModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderModel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderModel) ProtoMessage() {}

func (x *OrderModel) ProtoReflect() protoreflect.Message {
	mi := &file_orderModel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderModel.ProtoReflect.Descriptor instead.
func (*OrderModel) Descriptor() ([]byte, []int) {
	return file_orderModel_proto_rawDescGZIP(), []int{0}
}

func (x *OrderModel) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *OrderModel) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *OrderModel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *OrderModel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *OrderModel) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_orderModel_proto protoreflect.FileDescriptor

var file_orderModel_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0a,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_orderModel_proto_rawDescOnce sync.Once
	file_orderModel_proto_rawDescData = file_orderModel_proto_rawDesc
)

func file_orderModel_proto_rawDescGZIP() []byte {
	file_orderModel_proto_rawDescOnce.Do(func() {
		file_orderModel_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderModel_proto_rawDescData)
	})
	return file_orderModel_proto_rawDescData
}

var file_orderModel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_orderModel_proto_goTypes = []interface{}{
	(*OrderModel)(nil), // 0: service.OrderModel
}
var file_orderModel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orderModel_proto_init() }
func file_orderModel_proto_init() {
	if File_orderModel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderModel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderModel); i {
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
			RawDescriptor: file_orderModel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderModel_proto_goTypes,
		DependencyIndexes: file_orderModel_proto_depIdxs,
		MessageInfos:      file_orderModel_proto_msgTypes,
	}.Build()
	File_orderModel_proto = out.File
	file_orderModel_proto_rawDesc = nil
	file_orderModel_proto_goTypes = nil
	file_orderModel_proto_depIdxs = nil
}
