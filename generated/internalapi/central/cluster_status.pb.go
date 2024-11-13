// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: internalapi/central/cluster_status.proto

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

type DeploymentEnvironmentUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environments []string `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *DeploymentEnvironmentUpdate) Reset() {
	*x = DeploymentEnvironmentUpdate{}
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentEnvironmentUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentEnvironmentUpdate) ProtoMessage() {}

func (x *DeploymentEnvironmentUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentEnvironmentUpdate.ProtoReflect.Descriptor instead.
func (*DeploymentEnvironmentUpdate) Descriptor() ([]byte, []int) {
	return file_internalapi_central_cluster_status_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentEnvironmentUpdate) GetEnvironments() []string {
	if x != nil {
		return x.Environments
	}
	return nil
}

type ClusterStatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Making it a oneof for future proofing.
	//
	// Types that are assignable to Msg:
	//
	//	*ClusterStatusUpdate_Status
	//	*ClusterStatusUpdate_DeploymentEnvUpdate
	Msg isClusterStatusUpdate_Msg `protobuf_oneof:"msg"`
}

func (x *ClusterStatusUpdate) Reset() {
	*x = ClusterStatusUpdate{}
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatusUpdate) ProtoMessage() {}

func (x *ClusterStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatusUpdate.ProtoReflect.Descriptor instead.
func (*ClusterStatusUpdate) Descriptor() ([]byte, []int) {
	return file_internalapi_central_cluster_status_proto_rawDescGZIP(), []int{1}
}

func (m *ClusterStatusUpdate) GetMsg() isClusterStatusUpdate_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ClusterStatusUpdate) GetStatus() *storage.ClusterStatus {
	if x, ok := x.GetMsg().(*ClusterStatusUpdate_Status); ok {
		return x.Status
	}
	return nil
}

func (x *ClusterStatusUpdate) GetDeploymentEnvUpdate() *DeploymentEnvironmentUpdate {
	if x, ok := x.GetMsg().(*ClusterStatusUpdate_DeploymentEnvUpdate); ok {
		return x.DeploymentEnvUpdate
	}
	return nil
}

type isClusterStatusUpdate_Msg interface {
	isClusterStatusUpdate_Msg()
}

type ClusterStatusUpdate_Status struct {
	Status *storage.ClusterStatus `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type ClusterStatusUpdate_DeploymentEnvUpdate struct {
	DeploymentEnvUpdate *DeploymentEnvironmentUpdate `protobuf:"bytes,2,opt,name=deployment_env_update,json=deploymentEnvUpdate,proto3,oneof"`
}

func (*ClusterStatusUpdate_Status) isClusterStatusUpdate_Msg() {}

func (*ClusterStatusUpdate_DeploymentEnvUpdate) isClusterStatusUpdate_Msg() {}

type RawClusterHealthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectorHealthInfo        *storage.CollectorHealthInfo        `protobuf:"bytes,1,opt,name=collector_health_info,json=collectorHealthInfo,proto3" json:"collector_health_info,omitempty"`
	AdmissionControlHealthInfo *storage.AdmissionControlHealthInfo `protobuf:"bytes,2,opt,name=admission_control_health_info,json=admissionControlHealthInfo,proto3" json:"admission_control_health_info,omitempty"`
	ScannerHealthInfo          *storage.ScannerHealthInfo          `protobuf:"bytes,3,opt,name=scanner_health_info,json=scannerHealthInfo,proto3" json:"scanner_health_info,omitempty"`
}

func (x *RawClusterHealthInfo) Reset() {
	*x = RawClusterHealthInfo{}
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RawClusterHealthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawClusterHealthInfo) ProtoMessage() {}

func (x *RawClusterHealthInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawClusterHealthInfo.ProtoReflect.Descriptor instead.
func (*RawClusterHealthInfo) Descriptor() ([]byte, []int) {
	return file_internalapi_central_cluster_status_proto_rawDescGZIP(), []int{2}
}

func (x *RawClusterHealthInfo) GetCollectorHealthInfo() *storage.CollectorHealthInfo {
	if x != nil {
		return x.CollectorHealthInfo
	}
	return nil
}

func (x *RawClusterHealthInfo) GetAdmissionControlHealthInfo() *storage.AdmissionControlHealthInfo {
	if x != nil {
		return x.AdmissionControlHealthInfo
	}
	return nil
}

func (x *RawClusterHealthInfo) GetScannerHealthInfo() *storage.ScannerHealthInfo {
	if x != nil {
		return x.ScannerHealthInfo
	}
	return nil
}

type ClusterHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterHealthResponse) Reset() {
	*x = ClusterHealthResponse{}
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterHealthResponse) ProtoMessage() {}

func (x *ClusterHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_cluster_status_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterHealthResponse.ProtoReflect.Descriptor instead.
func (*ClusterHealthResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_cluster_status_proto_rawDescGZIP(), []int{3}
}

var File_internalapi_central_cluster_status_proto protoreflect.FileDescriptor

var file_internalapi_central_cluster_status_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x1a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x1b, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xaa, 0x01,
	0x0a, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5a, 0x0a, 0x15, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x13,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x9c, 0x02, 0x0a, 0x14, 0x52,
	0x61, 0x77, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x15, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x66, 0x0a, 0x1d, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x1a, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a,
	0x13, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x3b, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_central_cluster_status_proto_rawDescOnce sync.Once
	file_internalapi_central_cluster_status_proto_rawDescData = file_internalapi_central_cluster_status_proto_rawDesc
)

func file_internalapi_central_cluster_status_proto_rawDescGZIP() []byte {
	file_internalapi_central_cluster_status_proto_rawDescOnce.Do(func() {
		file_internalapi_central_cluster_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_central_cluster_status_proto_rawDescData)
	})
	return file_internalapi_central_cluster_status_proto_rawDescData
}

var file_internalapi_central_cluster_status_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internalapi_central_cluster_status_proto_goTypes = []any{
	(*DeploymentEnvironmentUpdate)(nil),        // 0: central.DeploymentEnvironmentUpdate
	(*ClusterStatusUpdate)(nil),                // 1: central.ClusterStatusUpdate
	(*RawClusterHealthInfo)(nil),               // 2: central.RawClusterHealthInfo
	(*ClusterHealthResponse)(nil),              // 3: central.ClusterHealthResponse
	(*storage.ClusterStatus)(nil),              // 4: storage.ClusterStatus
	(*storage.CollectorHealthInfo)(nil),        // 5: storage.CollectorHealthInfo
	(*storage.AdmissionControlHealthInfo)(nil), // 6: storage.AdmissionControlHealthInfo
	(*storage.ScannerHealthInfo)(nil),          // 7: storage.ScannerHealthInfo
}
var file_internalapi_central_cluster_status_proto_depIdxs = []int32{
	4, // 0: central.ClusterStatusUpdate.status:type_name -> storage.ClusterStatus
	0, // 1: central.ClusterStatusUpdate.deployment_env_update:type_name -> central.DeploymentEnvironmentUpdate
	5, // 2: central.RawClusterHealthInfo.collector_health_info:type_name -> storage.CollectorHealthInfo
	6, // 3: central.RawClusterHealthInfo.admission_control_health_info:type_name -> storage.AdmissionControlHealthInfo
	7, // 4: central.RawClusterHealthInfo.scanner_health_info:type_name -> storage.ScannerHealthInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internalapi_central_cluster_status_proto_init() }
func file_internalapi_central_cluster_status_proto_init() {
	if File_internalapi_central_cluster_status_proto != nil {
		return
	}
	file_internalapi_central_cluster_status_proto_msgTypes[1].OneofWrappers = []any{
		(*ClusterStatusUpdate_Status)(nil),
		(*ClusterStatusUpdate_DeploymentEnvUpdate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_central_cluster_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_cluster_status_proto_goTypes,
		DependencyIndexes: file_internalapi_central_cluster_status_proto_depIdxs,
		MessageInfos:      file_internalapi_central_cluster_status_proto_msgTypes,
	}.Build()
	File_internalapi_central_cluster_status_proto = out.File
	file_internalapi_central_cluster_status_proto_rawDesc = nil
	file_internalapi_central_cluster_status_proto_goTypes = nil
	file_internalapi_central_cluster_status_proto_depIdxs = nil
}
