// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: internalapi/central/auth.proto

package central

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceCertAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertDer     []byte                 `protobuf:"bytes,1,opt,name=cert_der,json=certDer,proto3" json:"cert_der,omitempty"`
	CurrentTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
}

func (x *ServiceCertAuth) Reset() {
	*x = ServiceCertAuth{}
	mi := &file_internalapi_central_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceCertAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceCertAuth) ProtoMessage() {}

func (x *ServiceCertAuth) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceCertAuth.ProtoReflect.Descriptor instead.
func (*ServiceCertAuth) Descriptor() ([]byte, []int) {
	return file_internalapi_central_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceCertAuth) GetCertDer() []byte {
	if x != nil {
		return x.CertDer
	}
	return nil
}

func (x *ServiceCertAuth) GetCurrentTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CurrentTime
	}
	return nil
}

var File_internalapi_central_auth_proto protoreflect.FileDescriptor

var file_internalapi_central_auth_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x65, 0x72, 0x74, 0x44, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x3b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_central_auth_proto_rawDescOnce sync.Once
	file_internalapi_central_auth_proto_rawDescData = file_internalapi_central_auth_proto_rawDesc
)

func file_internalapi_central_auth_proto_rawDescGZIP() []byte {
	file_internalapi_central_auth_proto_rawDescOnce.Do(func() {
		file_internalapi_central_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_central_auth_proto_rawDescData)
	})
	return file_internalapi_central_auth_proto_rawDescData
}

var file_internalapi_central_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internalapi_central_auth_proto_goTypes = []any{
	(*ServiceCertAuth)(nil),       // 0: central.ServiceCertAuth
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_internalapi_central_auth_proto_depIdxs = []int32{
	1, // 0: central.ServiceCertAuth.current_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_central_auth_proto_init() }
func file_internalapi_central_auth_proto_init() {
	if File_internalapi_central_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_central_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_auth_proto_goTypes,
		DependencyIndexes: file_internalapi_central_auth_proto_depIdxs,
		MessageInfos:      file_internalapi_central_auth_proto_msgTypes,
	}.Build()
	File_internalapi_central_auth_proto = out.File
	file_internalapi_central_auth_proto_rawDesc = nil
	file_internalapi_central_auth_proto_goTypes = nil
	file_internalapi_central_auth_proto_depIdxs = nil
}
