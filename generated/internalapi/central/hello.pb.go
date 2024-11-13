// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: internalapi/central/hello.proto

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

type SensorHello_SensorState int32

const (
	SensorHello_UNKNOWN   SensorHello_SensorState = 0
	SensorHello_STARTUP   SensorHello_SensorState = 1
	SensorHello_RECONNECT SensorHello_SensorState = 2
)

// Enum value maps for SensorHello_SensorState.
var (
	SensorHello_SensorState_name = map[int32]string{
		0: "UNKNOWN",
		1: "STARTUP",
		2: "RECONNECT",
	}
	SensorHello_SensorState_value = map[string]int32{
		"UNKNOWN":   0,
		"STARTUP":   1,
		"RECONNECT": 2,
	}
)

func (x SensorHello_SensorState) Enum() *SensorHello_SensorState {
	p := new(SensorHello_SensorState)
	*p = x
	return p
}

func (x SensorHello_SensorState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SensorHello_SensorState) Descriptor() protoreflect.EnumDescriptor {
	return file_internalapi_central_hello_proto_enumTypes[0].Descriptor()
}

func (SensorHello_SensorState) Type() protoreflect.EnumType {
	return &file_internalapi_central_hello_proto_enumTypes[0]
}

func (x SensorHello_SensorState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SensorHello_SensorState.Descriptor instead.
func (SensorHello_SensorState) EnumDescriptor() ([]byte, []int) {
	return file_internalapi_central_hello_proto_rawDescGZIP(), []int{1, 0}
}

type HelmManagedConfigInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterConfig *storage.CompleteClusterConfig `protobuf:"bytes,1,opt,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty"`
	ClusterName   string                         `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	ClusterId     string                         `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ManagedBy     storage.ManagerType            `protobuf:"varint,5,opt,name=managed_by,json=managedBy,proto3,enum=storage.ManagerType" json:"managed_by,omitempty"`
}

func (x *HelmManagedConfigInit) Reset() {
	*x = HelmManagedConfigInit{}
	mi := &file_internalapi_central_hello_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelmManagedConfigInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmManagedConfigInit) ProtoMessage() {}

func (x *HelmManagedConfigInit) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_hello_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmManagedConfigInit.ProtoReflect.Descriptor instead.
func (*HelmManagedConfigInit) Descriptor() ([]byte, []int) {
	return file_internalapi_central_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelmManagedConfigInit) GetClusterConfig() *storage.CompleteClusterConfig {
	if x != nil {
		return x.ClusterConfig
	}
	return nil
}

func (x *HelmManagedConfigInit) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *HelmManagedConfigInit) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *HelmManagedConfigInit) GetManagedBy() storage.ManagerType {
	if x != nil {
		return x.ManagedBy
	}
	return storage.ManagerType(0)
}

type SensorHello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorVersion            string                                  `protobuf:"bytes,1,opt,name=sensor_version,json=sensorVersion,proto3" json:"sensor_version,omitempty"`
	Capabilities             []string                                `protobuf:"bytes,2,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	DeploymentIdentification *storage.SensorDeploymentIdentification `protobuf:"bytes,5,opt,name=deployment_identification,json=deploymentIdentification,proto3" json:"deployment_identification,omitempty"`
	HelmManagedConfigInit    *HelmManagedConfigInit                  `protobuf:"bytes,3,opt,name=helm_managed_config_init,json=helmManagedConfigInit,proto3" json:"helm_managed_config_init,omitempty"`
	// Policy version sensor understands. If unset, central will try to guess it.
	PolicyVersion string `protobuf:"bytes,4,opt,name=policy_version,json=policyVersion,proto3" json:"policy_version,omitempty"`
	// sensor_state defines in what state sensor is (e.g. it's starting up or it's reconnecting)
	SensorState SensorHello_SensorState `protobuf:"varint,6,opt,name=sensor_state,json=sensorState,proto3,enum=central.SensorHello_SensorState" json:"sensor_state,omitempty"`
	// request_deduper_state is a request to central to communicate its deduper state.
	RequestDeduperState bool `protobuf:"varint,7,opt,name=request_deduper_state,json=requestDeduperState,proto3" json:"request_deduper_state,omitempty"`
}

func (x *SensorHello) Reset() {
	*x = SensorHello{}
	mi := &file_internalapi_central_hello_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SensorHello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorHello) ProtoMessage() {}

func (x *SensorHello) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_hello_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorHello.ProtoReflect.Descriptor instead.
func (*SensorHello) Descriptor() ([]byte, []int) {
	return file_internalapi_central_hello_proto_rawDescGZIP(), []int{1}
}

func (x *SensorHello) GetSensorVersion() string {
	if x != nil {
		return x.SensorVersion
	}
	return ""
}

func (x *SensorHello) GetCapabilities() []string {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *SensorHello) GetDeploymentIdentification() *storage.SensorDeploymentIdentification {
	if x != nil {
		return x.DeploymentIdentification
	}
	return nil
}

func (x *SensorHello) GetHelmManagedConfigInit() *HelmManagedConfigInit {
	if x != nil {
		return x.HelmManagedConfigInit
	}
	return nil
}

func (x *SensorHello) GetPolicyVersion() string {
	if x != nil {
		return x.PolicyVersion
	}
	return ""
}

func (x *SensorHello) GetSensorState() SensorHello_SensorState {
	if x != nil {
		return x.SensorState
	}
	return SensorHello_UNKNOWN
}

func (x *SensorHello) GetRequestDeduperState() bool {
	if x != nil {
		return x.RequestDeduperState
	}
	return false
}

type CentralHello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId        string            `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	CertBundle       map[string]string `protobuf:"bytes,2,rep,name=cert_bundle,json=certBundle,proto3" json:"cert_bundle,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ManagedCentral   bool              `protobuf:"varint,3,opt,name=managed_central,json=managedCentral,proto3" json:"managed_central,omitempty"`
	CentralId        string            `protobuf:"bytes,4,opt,name=central_id,json=centralId,proto3" json:"central_id,omitempty"`
	Capabilities     []string          `protobuf:"bytes,5,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	SendDeduperState bool              `protobuf:"varint,6,opt,name=send_deduper_state,json=sendDeduperState,proto3" json:"send_deduper_state,omitempty"`
}

func (x *CentralHello) Reset() {
	*x = CentralHello{}
	mi := &file_internalapi_central_hello_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CentralHello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CentralHello) ProtoMessage() {}

func (x *CentralHello) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_hello_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CentralHello.ProtoReflect.Descriptor instead.
func (*CentralHello) Descriptor() ([]byte, []int) {
	return file_internalapi_central_hello_proto_rawDescGZIP(), []int{2}
}

func (x *CentralHello) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CentralHello) GetCertBundle() map[string]string {
	if x != nil {
		return x.CertBundle
	}
	return nil
}

func (x *CentralHello) GetManagedCentral() bool {
	if x != nil {
		return x.ManagedCentral
	}
	return false
}

func (x *CentralHello) GetCentralId() string {
	if x != nil {
		return x.CentralId
	}
	return ""
}

func (x *CentralHello) GetCapabilities() []string {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *CentralHello) GetSendDeduperState() bool {
	if x != nil {
		return x.SendDeduperState
	}
	return false
}

var File_internalapi_central_hello_proto protoreflect.FileDescriptor

var file_internalapi_central_hello_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x1a, 0x15, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x15, 0x48, 0x65, 0x6c, 0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x45, 0x0a, 0x0e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x42, 0x79, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22,
	0xef, 0x03, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x19, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x57, 0x0a, 0x18, 0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x48, 0x65, 0x6c,
	0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e,
	0x69, 0x74, 0x52, 0x15, 0x68, 0x65, 0x6c, 0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x43, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x64, 0x65, 0x64, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x64,
	0x75, 0x70, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x55, 0x50,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10,
	0x02, 0x22, 0xce, 0x02, 0x0a, 0x0c, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x46, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x2e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x43, 0x65,
	0x72, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63,
	0x65, 0x72, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x43, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65,
	0x64, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x64, 0x75, 0x70, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x65, 0x72, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x3b, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_central_hello_proto_rawDescOnce sync.Once
	file_internalapi_central_hello_proto_rawDescData = file_internalapi_central_hello_proto_rawDesc
)

func file_internalapi_central_hello_proto_rawDescGZIP() []byte {
	file_internalapi_central_hello_proto_rawDescOnce.Do(func() {
		file_internalapi_central_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_central_hello_proto_rawDescData)
	})
	return file_internalapi_central_hello_proto_rawDescData
}

var file_internalapi_central_hello_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internalapi_central_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internalapi_central_hello_proto_goTypes = []any{
	(SensorHello_SensorState)(0),                   // 0: central.SensorHello.SensorState
	(*HelmManagedConfigInit)(nil),                  // 1: central.HelmManagedConfigInit
	(*SensorHello)(nil),                            // 2: central.SensorHello
	(*CentralHello)(nil),                           // 3: central.CentralHello
	nil,                                            // 4: central.CentralHello.CertBundleEntry
	(*storage.CompleteClusterConfig)(nil),          // 5: storage.CompleteClusterConfig
	(storage.ManagerType)(0),                       // 6: storage.ManagerType
	(*storage.SensorDeploymentIdentification)(nil), // 7: storage.SensorDeploymentIdentification
}
var file_internalapi_central_hello_proto_depIdxs = []int32{
	5, // 0: central.HelmManagedConfigInit.cluster_config:type_name -> storage.CompleteClusterConfig
	6, // 1: central.HelmManagedConfigInit.managed_by:type_name -> storage.ManagerType
	7, // 2: central.SensorHello.deployment_identification:type_name -> storage.SensorDeploymentIdentification
	1, // 3: central.SensorHello.helm_managed_config_init:type_name -> central.HelmManagedConfigInit
	0, // 4: central.SensorHello.sensor_state:type_name -> central.SensorHello.SensorState
	4, // 5: central.CentralHello.cert_bundle:type_name -> central.CentralHello.CertBundleEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_internalapi_central_hello_proto_init() }
func file_internalapi_central_hello_proto_init() {
	if File_internalapi_central_hello_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_central_hello_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_hello_proto_goTypes,
		DependencyIndexes: file_internalapi_central_hello_proto_depIdxs,
		EnumInfos:         file_internalapi_central_hello_proto_enumTypes,
		MessageInfos:      file_internalapi_central_hello_proto_msgTypes,
	}.Build()
	File_internalapi_central_hello_proto = out.File
	file_internalapi_central_hello_proto_rawDesc = nil
	file_internalapi_central_hello_proto_goTypes = nil
	file_internalapi_central_hello_proto_depIdxs = nil
}
