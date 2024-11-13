// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: api/v1/notifications.proto

package v1

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

type NetworkPolicyNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Yaml    string `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *NetworkPolicyNotification) Reset() {
	*x = NetworkPolicyNotification{}
	mi := &file_api_v1_notifications_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkPolicyNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkPolicyNotification) ProtoMessage() {}

func (x *NetworkPolicyNotification) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_notifications_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkPolicyNotification.ProtoReflect.Descriptor instead.
func (*NetworkPolicyNotification) Descriptor() ([]byte, []int) {
	return file_api_v1_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkPolicyNotification) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *NetworkPolicyNotification) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

var File_api_v1_notifications_proto protoreflect.FileDescriptor

var file_api_v1_notifications_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x22, 0x49, 0x0a, 0x19, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x42, 0x27, 0x0a, 0x18, 0x69,
	0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_notifications_proto_rawDescOnce sync.Once
	file_api_v1_notifications_proto_rawDescData = file_api_v1_notifications_proto_rawDesc
)

func file_api_v1_notifications_proto_rawDescGZIP() []byte {
	file_api_v1_notifications_proto_rawDescOnce.Do(func() {
		file_api_v1_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_notifications_proto_rawDescData)
	})
	return file_api_v1_notifications_proto_rawDescData
}

var file_api_v1_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_notifications_proto_goTypes = []any{
	(*NetworkPolicyNotification)(nil), // 0: v1.NetworkPolicyNotification
}
var file_api_v1_notifications_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_notifications_proto_init() }
func file_api_v1_notifications_proto_init() {
	if File_api_v1_notifications_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_notifications_proto_goTypes,
		DependencyIndexes: file_api_v1_notifications_proto_depIdxs,
		MessageInfos:      file_api_v1_notifications_proto_msgTypes,
	}.Build()
	File_api_v1_notifications_proto = out.File
	file_api_v1_notifications_proto_rawDesc = nil
	file_api_v1_notifications_proto_goTypes = nil
	file_api_v1_notifications_proto_depIdxs = nil
}
