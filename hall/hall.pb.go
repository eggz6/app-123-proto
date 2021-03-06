// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: hall.proto

package hall

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

type RegisteryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegisteryRequest) Reset() {
	*x = RegisteryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteryRequest) ProtoMessage() {}

func (x *RegisteryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteryRequest.ProtoReflect.Descriptor instead.
func (*RegisteryRequest) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{0}
}

func (x *RegisteryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisteryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RegisteryResponse) Reset() {
	*x = RegisteryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteryResponse) ProtoMessage() {}

func (x *RegisteryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteryResponse.ProtoReflect.Descriptor instead.
func (*RegisteryResponse) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{1}
}

func (x *RegisteryResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_hall_proto protoreflect.FileDescriptor

var file_hall_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x61,
	0x6c, 0x6c, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x46, 0x0a, 0x09, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x79, 0x12, 0x39,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x68, 0x61, 0x6c, 0x6c, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x68, 0x61, 0x6c, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b,
	0x68, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hall_proto_rawDescOnce sync.Once
	file_hall_proto_rawDescData = file_hall_proto_rawDesc
)

func file_hall_proto_rawDescGZIP() []byte {
	file_hall_proto_rawDescOnce.Do(func() {
		file_hall_proto_rawDescData = protoimpl.X.CompressGZIP(file_hall_proto_rawDescData)
	})
	return file_hall_proto_rawDescData
}

var file_hall_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hall_proto_goTypes = []interface{}{
	(*RegisteryRequest)(nil),  // 0: hall.RegisteryRequest
	(*RegisteryResponse)(nil), // 1: hall.RegisteryResponse
}
var file_hall_proto_depIdxs = []int32{
	0, // 0: hall.Registery.Ping:input_type -> hall.RegisteryRequest
	1, // 1: hall.Registery.Ping:output_type -> hall.RegisteryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hall_proto_init() }
func file_hall_proto_init() {
	if File_hall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteryRequest); i {
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
		file_hall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteryResponse); i {
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
			RawDescriptor: file_hall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hall_proto_goTypes,
		DependencyIndexes: file_hall_proto_depIdxs,
		MessageInfos:      file_hall_proto_msgTypes,
	}.Build()
	File_hall_proto = out.File
	file_hall_proto_rawDesc = nil
	file_hall_proto_goTypes = nil
	file_hall_proto_depIdxs = nil
}
