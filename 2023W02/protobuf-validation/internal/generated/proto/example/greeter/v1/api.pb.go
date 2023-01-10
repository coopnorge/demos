// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: example/greeter/v1/api.proto

package greeterv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "gitlab.com/aucampia/eg/service-golang/internal/generated/proto/example/types/v1"
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

// The request message containing the name of the entity to be greeted.
type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SomeValue *v1.SomeType `protobuf:"bytes,2,opt,name=some_value,json=someValue,proto3" json:"some_value,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_greeter_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_greeter_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_example_greeter_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GreetRequest) GetSomeValue() *v1.SomeType {
	if x != nil {
		return x.SomeValue
	}
	return nil
}

// The response message containing the greetings.
type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_greeter_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_greeter_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_example_greeter_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_example_greeter_v1_api_proto protoreflect.FileDescriptor

var file_example_greeter_v1_api_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xfa, 0x42, 0x1e,
	0x72, 0x1c, 0x10, 0x01, 0x5a, 0x0a, 0x4a, 0x61, 0x63, 0x6b, 0x20, 0x53, 0x6d, 0x69, 0x74, 0x68,
	0x5a, 0x0c, 0x4a, 0x61, 0x63, 0x6b, 0x69, 0x74, 0x79, 0x20, 0x4a, 0x61, 0x63, 0x6b, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09,
	0x73, 0x6f, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x5a, 0x0a, 0x0a, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x41,
	0x50, 0x49, 0x12, 0x4c, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xe9, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x63, 0x61, 0x6d, 0x70, 0x69, 0x61, 0x2f, 0x65, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x47, 0x58, 0xaa, 0x02, 0x12, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x12, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a,
	0x3a, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_greeter_v1_api_proto_rawDescOnce sync.Once
	file_example_greeter_v1_api_proto_rawDescData = file_example_greeter_v1_api_proto_rawDesc
)

func file_example_greeter_v1_api_proto_rawDescGZIP() []byte {
	file_example_greeter_v1_api_proto_rawDescOnce.Do(func() {
		file_example_greeter_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_greeter_v1_api_proto_rawDescData)
	})
	return file_example_greeter_v1_api_proto_rawDescData
}

var file_example_greeter_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_greeter_v1_api_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),  // 0: example.greeter.v1.GreetRequest
	(*GreetResponse)(nil), // 1: example.greeter.v1.GreetResponse
	(*v1.SomeType)(nil),   // 2: example.types.v1.SomeType
}
var file_example_greeter_v1_api_proto_depIdxs = []int32{
	2, // 0: example.greeter.v1.GreetRequest.some_value:type_name -> example.types.v1.SomeType
	0, // 1: example.greeter.v1.GreeterAPI.Greet:input_type -> example.greeter.v1.GreetRequest
	1, // 2: example.greeter.v1.GreeterAPI.Greet:output_type -> example.greeter.v1.GreetResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_greeter_v1_api_proto_init() }
func file_example_greeter_v1_api_proto_init() {
	if File_example_greeter_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_greeter_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_example_greeter_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
			RawDescriptor: file_example_greeter_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_greeter_v1_api_proto_goTypes,
		DependencyIndexes: file_example_greeter_v1_api_proto_depIdxs,
		MessageInfos:      file_example_greeter_v1_api_proto_msgTypes,
	}.Build()
	File_example_greeter_v1_api_proto = out.File
	file_example_greeter_v1_api_proto_rawDesc = nil
	file_example_greeter_v1_api_proto_goTypes = nil
	file_example_greeter_v1_api_proto_depIdxs = nil
}
