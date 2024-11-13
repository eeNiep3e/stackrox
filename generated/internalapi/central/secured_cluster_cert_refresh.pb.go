// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: internalapi/central/secured_cluster_cert_refresh.proto

package central

import (
	storage "github.com/stackrox/rox/generated/storage"
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

type SecuredClusterCertsIssueError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SecuredClusterCertsIssueError) Reset() {
	*x = SecuredClusterCertsIssueError{}
	mi := &file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecuredClusterCertsIssueError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecuredClusterCertsIssueError) ProtoMessage() {}

func (x *SecuredClusterCertsIssueError) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecuredClusterCertsIssueError.ProtoReflect.Descriptor instead.
func (*SecuredClusterCertsIssueError) Descriptor() ([]byte, []int) {
	return file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescGZIP(), []int{0}
}

func (x *SecuredClusterCertsIssueError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IssueSecuredClusterCertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *IssueSecuredClusterCertsRequest) Reset() {
	*x = IssueSecuredClusterCertsRequest{}
	mi := &file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueSecuredClusterCertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueSecuredClusterCertsRequest) ProtoMessage() {}

func (x *IssueSecuredClusterCertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueSecuredClusterCertsRequest.ProtoReflect.Descriptor instead.
func (*IssueSecuredClusterCertsRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescGZIP(), []int{1}
}

func (x *IssueSecuredClusterCertsRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type IssueSecuredClusterCertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are assignable to Response:
	//
	//	*IssueSecuredClusterCertsResponse_Certificates
	//	*IssueSecuredClusterCertsResponse_Error
	Response isIssueSecuredClusterCertsResponse_Response `protobuf_oneof:"response"`
}

func (x *IssueSecuredClusterCertsResponse) Reset() {
	*x = IssueSecuredClusterCertsResponse{}
	mi := &file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueSecuredClusterCertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueSecuredClusterCertsResponse) ProtoMessage() {}

func (x *IssueSecuredClusterCertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueSecuredClusterCertsResponse.ProtoReflect.Descriptor instead.
func (*IssueSecuredClusterCertsResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescGZIP(), []int{2}
}

func (x *IssueSecuredClusterCertsResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (m *IssueSecuredClusterCertsResponse) GetResponse() isIssueSecuredClusterCertsResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *IssueSecuredClusterCertsResponse) GetCertificates() *storage.TypedServiceCertificateSet {
	if x, ok := x.GetResponse().(*IssueSecuredClusterCertsResponse_Certificates); ok {
		return x.Certificates
	}
	return nil
}

func (x *IssueSecuredClusterCertsResponse) GetError() *SecuredClusterCertsIssueError {
	if x, ok := x.GetResponse().(*IssueSecuredClusterCertsResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isIssueSecuredClusterCertsResponse_Response interface {
	isIssueSecuredClusterCertsResponse_Response()
}

type IssueSecuredClusterCertsResponse_Certificates struct {
	Certificates *storage.TypedServiceCertificateSet `protobuf:"bytes,2,opt,name=certificates,proto3,oneof"`
}

type IssueSecuredClusterCertsResponse_Error struct {
	Error *SecuredClusterCertsIssueError `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*IssueSecuredClusterCertsResponse_Certificates) isIssueSecuredClusterCertsResponse_Response() {}

func (*IssueSecuredClusterCertsResponse_Error) isIssueSecuredClusterCertsResponse_Response() {}

var File_internalapi_central_secured_cluster_cert_refresh_proto protoreflect.FileDescriptor

var file_internalapi_central_secured_cluster_cert_refresh_proto_rawDesc = []byte{
	0x0a, 0x36, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x1a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x39, 0x0a, 0x1d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x1f,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xd8,
	0x01, 0x0a, 0x20, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x49, 0x0a, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52,
	0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x3e, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x3b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescOnce sync.Once
	file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescData = file_internalapi_central_secured_cluster_cert_refresh_proto_rawDesc
)

func file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescGZIP() []byte {
	file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescOnce.Do(func() {
		file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescData)
	})
	return file_internalapi_central_secured_cluster_cert_refresh_proto_rawDescData
}

var file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internalapi_central_secured_cluster_cert_refresh_proto_goTypes = []any{
	(*SecuredClusterCertsIssueError)(nil),      // 0: central.SecuredClusterCertsIssueError
	(*IssueSecuredClusterCertsRequest)(nil),    // 1: central.IssueSecuredClusterCertsRequest
	(*IssueSecuredClusterCertsResponse)(nil),   // 2: central.IssueSecuredClusterCertsResponse
	(*storage.TypedServiceCertificateSet)(nil), // 3: storage.TypedServiceCertificateSet
}
var file_internalapi_central_secured_cluster_cert_refresh_proto_depIdxs = []int32{
	3, // 0: central.IssueSecuredClusterCertsResponse.certificates:type_name -> storage.TypedServiceCertificateSet
	0, // 1: central.IssueSecuredClusterCertsResponse.error:type_name -> central.SecuredClusterCertsIssueError
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internalapi_central_secured_cluster_cert_refresh_proto_init() }
func file_internalapi_central_secured_cluster_cert_refresh_proto_init() {
	if File_internalapi_central_secured_cluster_cert_refresh_proto != nil {
		return
	}
	file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes[2].OneofWrappers = []any{
		(*IssueSecuredClusterCertsResponse_Certificates)(nil),
		(*IssueSecuredClusterCertsResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_central_secured_cluster_cert_refresh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_secured_cluster_cert_refresh_proto_goTypes,
		DependencyIndexes: file_internalapi_central_secured_cluster_cert_refresh_proto_depIdxs,
		MessageInfos:      file_internalapi_central_secured_cluster_cert_refresh_proto_msgTypes,
	}.Build()
	File_internalapi_central_secured_cluster_cert_refresh_proto = out.File
	file_internalapi_central_secured_cluster_cert_refresh_proto_rawDesc = nil
	file_internalapi_central_secured_cluster_cert_refresh_proto_goTypes = nil
	file_internalapi_central_secured_cluster_cert_refresh_proto_depIdxs = nil
}
