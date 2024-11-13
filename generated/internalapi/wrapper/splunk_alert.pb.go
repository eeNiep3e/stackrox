// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: internalapi/wrapper/splunk_alert.proto

package wrapper

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Splunk notification needs the source of data
// and the type of data.
type SplunkEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event      *anypb.Any `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Source     string     `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Sourcetype string     `protobuf:"bytes,3,opt,name=sourcetype,proto3" json:"sourcetype,omitempty"`
}

func (x *SplunkEvent) Reset() {
	*x = SplunkEvent{}
	mi := &file_internalapi_wrapper_splunk_alert_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SplunkEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplunkEvent) ProtoMessage() {}

func (x *SplunkEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_wrapper_splunk_alert_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplunkEvent.ProtoReflect.Descriptor instead.
func (*SplunkEvent) Descriptor() ([]byte, []int) {
	return file_internalapi_wrapper_splunk_alert_proto_rawDescGZIP(), []int{0}
}

func (x *SplunkEvent) GetEvent() *anypb.Any {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *SplunkEvent) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SplunkEvent) GetSourcetype() string {
	if x != nil {
		return x.Sourcetype
	}
	return ""
}

var File_internalapi_wrapper_splunk_alert_proto protoreflect.FileDescriptor

var file_internalapi_wrapper_splunk_alert_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x70, 0x6c, 0x75, 0x6e, 0x6b, 0x5f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0b,
	0x53, 0x70, 0x6c, 0x75, 0x6e, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x3b, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_wrapper_splunk_alert_proto_rawDescOnce sync.Once
	file_internalapi_wrapper_splunk_alert_proto_rawDescData = file_internalapi_wrapper_splunk_alert_proto_rawDesc
)

func file_internalapi_wrapper_splunk_alert_proto_rawDescGZIP() []byte {
	file_internalapi_wrapper_splunk_alert_proto_rawDescOnce.Do(func() {
		file_internalapi_wrapper_splunk_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_wrapper_splunk_alert_proto_rawDescData)
	})
	return file_internalapi_wrapper_splunk_alert_proto_rawDescData
}

var file_internalapi_wrapper_splunk_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internalapi_wrapper_splunk_alert_proto_goTypes = []any{
	(*SplunkEvent)(nil), // 0: wrapper.SplunkEvent
	(*anypb.Any)(nil),   // 1: google.protobuf.Any
}
var file_internalapi_wrapper_splunk_alert_proto_depIdxs = []int32{
	1, // 0: wrapper.SplunkEvent.event:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_wrapper_splunk_alert_proto_init() }
func file_internalapi_wrapper_splunk_alert_proto_init() {
	if File_internalapi_wrapper_splunk_alert_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_wrapper_splunk_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_wrapper_splunk_alert_proto_goTypes,
		DependencyIndexes: file_internalapi_wrapper_splunk_alert_proto_depIdxs,
		MessageInfos:      file_internalapi_wrapper_splunk_alert_proto_msgTypes,
	}.Build()
	File_internalapi_wrapper_splunk_alert_proto = out.File
	file_internalapi_wrapper_splunk_alert_proto_rawDesc = nil
	file_internalapi_wrapper_splunk_alert_proto_goTypes = nil
	file_internalapi_wrapper_splunk_alert_proto_depIdxs = nil
}
