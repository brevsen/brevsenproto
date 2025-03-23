// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: distribution/v1/distribution.proto

package distributionv1

import (
	v1 "github.com/brevsen/brevsenproto/gen/go/bp/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FetchOperations
type FetchOperationsRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OriginSnapshot uint32                 `protobuf:"varint,1,opt,name=origin_snapshot,json=originSnapshot,proto3" json:"origin_snapshot,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FetchOperationsRequest) Reset() {
	*x = FetchOperationsRequest{}
	mi := &file_distribution_v1_distribution_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchOperationsRequest) ProtoMessage() {}

func (x *FetchOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_distribution_v1_distribution_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchOperationsRequest.ProtoReflect.Descriptor instead.
func (*FetchOperationsRequest) Descriptor() ([]byte, []int) {
	return file_distribution_v1_distribution_proto_rawDescGZIP(), []int{0}
}

func (x *FetchOperationsRequest) GetOriginSnapshot() uint32 {
	if x != nil {
		return x.OriginSnapshot
	}
	return 0
}

type FetchOperationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operation     *v1.Operation          `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchOperationsResponse) Reset() {
	*x = FetchOperationsResponse{}
	mi := &file_distribution_v1_distribution_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchOperationsResponse) ProtoMessage() {}

func (x *FetchOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_distribution_v1_distribution_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchOperationsResponse.ProtoReflect.Descriptor instead.
func (*FetchOperationsResponse) Descriptor() ([]byte, []int) {
	return file_distribution_v1_distribution_proto_rawDescGZIP(), []int{1}
}

func (x *FetchOperationsResponse) GetOperation() *v1.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

var File_distribution_v1_distribution_proto protoreflect.FileDescriptor

var file_distribution_v1_distribution_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x22, 0x4d, 0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x7d, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0xcf,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x65, 0x76, 0x73, 0x65, 0x6e, 0x2f,
	0x62, 0x72, 0x65, 0x76, 0x73, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1b, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_distribution_v1_distribution_proto_rawDescOnce sync.Once
	file_distribution_v1_distribution_proto_rawDescData []byte
)

func file_distribution_v1_distribution_proto_rawDescGZIP() []byte {
	file_distribution_v1_distribution_proto_rawDescOnce.Do(func() {
		file_distribution_v1_distribution_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_distribution_v1_distribution_proto_rawDesc), len(file_distribution_v1_distribution_proto_rawDesc)))
	})
	return file_distribution_v1_distribution_proto_rawDescData
}

var file_distribution_v1_distribution_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_distribution_v1_distribution_proto_goTypes = []any{
	(*FetchOperationsRequest)(nil),  // 0: distribution.v1.FetchOperationsRequest
	(*FetchOperationsResponse)(nil), // 1: distribution.v1.FetchOperationsResponse
	(*v1.Operation)(nil),            // 2: common.v1.Operation
}
var file_distribution_v1_distribution_proto_depIdxs = []int32{
	2, // 0: distribution.v1.FetchOperationsResponse.operation:type_name -> common.v1.Operation
	0, // 1: distribution.v1.DistributionService.FetchOperations:input_type -> distribution.v1.FetchOperationsRequest
	1, // 2: distribution.v1.DistributionService.FetchOperations:output_type -> distribution.v1.FetchOperationsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_distribution_v1_distribution_proto_init() }
func file_distribution_v1_distribution_proto_init() {
	if File_distribution_v1_distribution_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_distribution_v1_distribution_proto_rawDesc), len(file_distribution_v1_distribution_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_distribution_v1_distribution_proto_goTypes,
		DependencyIndexes: file_distribution_v1_distribution_proto_depIdxs,
		MessageInfos:      file_distribution_v1_distribution_proto_msgTypes,
	}.Build()
	File_distribution_v1_distribution_proto = out.File
	file_distribution_v1_distribution_proto_goTypes = nil
	file_distribution_v1_distribution_proto_depIdxs = nil
}
