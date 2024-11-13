// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: api/v1/credential_expiry_service.proto

package v1

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

type GetCertExpiry_Component int32

const (
	GetCertExpiry_UNKNOWN    GetCertExpiry_Component = 0
	GetCertExpiry_CENTRAL    GetCertExpiry_Component = 1
	GetCertExpiry_SCANNER    GetCertExpiry_Component = 2
	GetCertExpiry_SCANNER_V4 GetCertExpiry_Component = 3
)

// Enum value maps for GetCertExpiry_Component.
var (
	GetCertExpiry_Component_name = map[int32]string{
		0: "UNKNOWN",
		1: "CENTRAL",
		2: "SCANNER",
		3: "SCANNER_V4",
	}
	GetCertExpiry_Component_value = map[string]int32{
		"UNKNOWN":    0,
		"CENTRAL":    1,
		"SCANNER":    2,
		"SCANNER_V4": 3,
	}
)

func (x GetCertExpiry_Component) Enum() *GetCertExpiry_Component {
	p := new(GetCertExpiry_Component)
	*p = x
	return p
}

func (x GetCertExpiry_Component) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetCertExpiry_Component) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_credential_expiry_service_proto_enumTypes[0].Descriptor()
}

func (GetCertExpiry_Component) Type() protoreflect.EnumType {
	return &file_api_v1_credential_expiry_service_proto_enumTypes[0]
}

func (x GetCertExpiry_Component) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetCertExpiry_Component.Descriptor instead.
func (GetCertExpiry_Component) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_credential_expiry_service_proto_rawDescGZIP(), []int{0, 0}
}

type GetCertExpiry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCertExpiry) Reset() {
	*x = GetCertExpiry{}
	mi := &file_api_v1_credential_expiry_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCertExpiry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertExpiry) ProtoMessage() {}

func (x *GetCertExpiry) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_credential_expiry_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertExpiry.ProtoReflect.Descriptor instead.
func (*GetCertExpiry) Descriptor() ([]byte, []int) {
	return file_api_v1_credential_expiry_service_proto_rawDescGZIP(), []int{0}
}

type GetCertExpiry_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component GetCertExpiry_Component `protobuf:"varint,1,opt,name=component,proto3,enum=v1.GetCertExpiry_Component" json:"component,omitempty"`
}

func (x *GetCertExpiry_Request) Reset() {
	*x = GetCertExpiry_Request{}
	mi := &file_api_v1_credential_expiry_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCertExpiry_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertExpiry_Request) ProtoMessage() {}

func (x *GetCertExpiry_Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_credential_expiry_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertExpiry_Request.ProtoReflect.Descriptor instead.
func (*GetCertExpiry_Request) Descriptor() ([]byte, []int) {
	return file_api_v1_credential_expiry_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetCertExpiry_Request) GetComponent() GetCertExpiry_Component {
	if x != nil {
		return x.Component
	}
	return GetCertExpiry_UNKNOWN
}

type GetCertExpiry_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expiry *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *GetCertExpiry_Response) Reset() {
	*x = GetCertExpiry_Response{}
	mi := &file_api_v1_credential_expiry_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCertExpiry_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertExpiry_Response) ProtoMessage() {}

func (x *GetCertExpiry_Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_credential_expiry_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertExpiry_Response.ProtoReflect.Descriptor instead.
func (*GetCertExpiry_Response) Descriptor() ([]byte, []int) {
	return file_api_v1_credential_expiry_service_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GetCertExpiry_Response) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

var File_api_v1_credential_expiry_service_proto protoreflect.FileDescriptor

var file_api_v1_credential_expiry_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x1a, 0x44, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x1a, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x43,
	0x41, 0x4e, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x41, 0x4e, 0x4e,
	0x45, 0x52, 0x5f, 0x56, 0x34, 0x10, 0x03, 0x32, 0x7f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x58, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_credential_expiry_service_proto_rawDescOnce sync.Once
	file_api_v1_credential_expiry_service_proto_rawDescData = file_api_v1_credential_expiry_service_proto_rawDesc
)

func file_api_v1_credential_expiry_service_proto_rawDescGZIP() []byte {
	file_api_v1_credential_expiry_service_proto_rawDescOnce.Do(func() {
		file_api_v1_credential_expiry_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_credential_expiry_service_proto_rawDescData)
	})
	return file_api_v1_credential_expiry_service_proto_rawDescData
}

var file_api_v1_credential_expiry_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_credential_expiry_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_credential_expiry_service_proto_goTypes = []any{
	(GetCertExpiry_Component)(0),   // 0: v1.GetCertExpiry.Component
	(*GetCertExpiry)(nil),          // 1: v1.GetCertExpiry
	(*GetCertExpiry_Request)(nil),  // 2: v1.GetCertExpiry.Request
	(*GetCertExpiry_Response)(nil), // 3: v1.GetCertExpiry.Response
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_api_v1_credential_expiry_service_proto_depIdxs = []int32{
	0, // 0: v1.GetCertExpiry.Request.component:type_name -> v1.GetCertExpiry.Component
	4, // 1: v1.GetCertExpiry.Response.expiry:type_name -> google.protobuf.Timestamp
	2, // 2: v1.CredentialExpiryService.GetCertExpiry:input_type -> v1.GetCertExpiry.Request
	3, // 3: v1.CredentialExpiryService.GetCertExpiry:output_type -> v1.GetCertExpiry.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_credential_expiry_service_proto_init() }
func file_api_v1_credential_expiry_service_proto_init() {
	if File_api_v1_credential_expiry_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_credential_expiry_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_credential_expiry_service_proto_goTypes,
		DependencyIndexes: file_api_v1_credential_expiry_service_proto_depIdxs,
		EnumInfos:         file_api_v1_credential_expiry_service_proto_enumTypes,
		MessageInfos:      file_api_v1_credential_expiry_service_proto_msgTypes,
	}.Build()
	File_api_v1_credential_expiry_service_proto = out.File
	file_api_v1_credential_expiry_service_proto_rawDesc = nil
	file_api_v1_credential_expiry_service_proto_goTypes = nil
	file_api_v1_credential_expiry_service_proto_depIdxs = nil
}
