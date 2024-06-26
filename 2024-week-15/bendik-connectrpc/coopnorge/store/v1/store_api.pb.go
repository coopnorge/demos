// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: coopnorge/store/v1/store_api.proto

package storev1

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

// Parameters for GetStore request.
type GetStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier for the Store.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStoreRequest) Reset() {
	*x = GetStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coopnorge_store_v1_store_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoreRequest) ProtoMessage() {}

func (x *GetStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coopnorge_store_v1_store_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoreRequest.ProtoReflect.Descriptor instead.
func (*GetStoreRequest) Descriptor() ([]byte, []int) {
	return file_coopnorge_store_v1_store_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetStoreRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response for GetStore
type GetStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Store *Store `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
}

func (x *GetStoreResponse) Reset() {
	*x = GetStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coopnorge_store_v1_store_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoreResponse) ProtoMessage() {}

func (x *GetStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coopnorge_store_v1_store_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoreResponse.ProtoReflect.Descriptor instead.
func (*GetStoreResponse) Descriptor() ([]byte, []int) {
	return file_coopnorge_store_v1_store_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetStoreResponse) GetStore() *Store {
	if x != nil {
		return x.Store
	}
	return nil
}

// The message describes a Store which is a physical outlet which is part of a
// specific coop chain.
type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier for the Store.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The store's name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coopnorge_store_v1_store_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_coopnorge_store_v1_store_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_coopnorge_store_v1_store_api_proto_rawDescGZIP(), []int{2}
}

func (x *Store) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_coopnorge_store_v1_store_api_proto protoreflect.FileDescriptor

var file_coopnorge_store_v1_store_api_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x22, 0x2b, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x6c, 0x0a,
	0x13, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x50, 0x49, 0x12, 0x55, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd4, 0x01, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x62, 0x65,
	0x6e, 0x64, 0x69, 0x6b, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58,
	0xaa, 0x02, 0x12, 0x43, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x6f, 0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67,
	0x65, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x6f, 0x6f,
	0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x6f,
	0x6f, 0x70, 0x6e, 0x6f, 0x72, 0x67, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coopnorge_store_v1_store_api_proto_rawDescOnce sync.Once
	file_coopnorge_store_v1_store_api_proto_rawDescData = file_coopnorge_store_v1_store_api_proto_rawDesc
)

func file_coopnorge_store_v1_store_api_proto_rawDescGZIP() []byte {
	file_coopnorge_store_v1_store_api_proto_rawDescOnce.Do(func() {
		file_coopnorge_store_v1_store_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_coopnorge_store_v1_store_api_proto_rawDescData)
	})
	return file_coopnorge_store_v1_store_api_proto_rawDescData
}

var file_coopnorge_store_v1_store_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coopnorge_store_v1_store_api_proto_goTypes = []interface{}{
	(*GetStoreRequest)(nil),  // 0: coopnorge.store.v1.GetStoreRequest
	(*GetStoreResponse)(nil), // 1: coopnorge.store.v1.GetStoreResponse
	(*Store)(nil),            // 2: coopnorge.store.v1.Store
}
var file_coopnorge_store_v1_store_api_proto_depIdxs = []int32{
	2, // 0: coopnorge.store.v1.GetStoreResponse.store:type_name -> coopnorge.store.v1.Store
	0, // 1: coopnorge.store.v1.StoreInformationAPI.GetStore:input_type -> coopnorge.store.v1.GetStoreRequest
	1, // 2: coopnorge.store.v1.StoreInformationAPI.GetStore:output_type -> coopnorge.store.v1.GetStoreResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coopnorge_store_v1_store_api_proto_init() }
func file_coopnorge_store_v1_store_api_proto_init() {
	if File_coopnorge_store_v1_store_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coopnorge_store_v1_store_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoreRequest); i {
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
		file_coopnorge_store_v1_store_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoreResponse); i {
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
		file_coopnorge_store_v1_store_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
			RawDescriptor: file_coopnorge_store_v1_store_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coopnorge_store_v1_store_api_proto_goTypes,
		DependencyIndexes: file_coopnorge_store_v1_store_api_proto_depIdxs,
		MessageInfos:      file_coopnorge_store_v1_store_api_proto_msgTypes,
	}.Build()
	File_coopnorge_store_v1_store_api_proto = out.File
	file_coopnorge_store_v1_store_api_proto_rawDesc = nil
	file_coopnorge_store_v1_store_api_proto_goTypes = nil
	file_coopnorge_store_v1_store_api_proto_depIdxs = nil
}
