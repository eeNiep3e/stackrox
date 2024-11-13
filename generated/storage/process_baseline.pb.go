// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: storage/process_baseline.proto

package storage

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

type ContainerNameAndBaselineStatus_BaselineStatus int32

const (
	ContainerNameAndBaselineStatus_INVALID       ContainerNameAndBaselineStatus_BaselineStatus = 0
	ContainerNameAndBaselineStatus_NOT_GENERATED ContainerNameAndBaselineStatus_BaselineStatus = 1 // In current implementation, this is a temporary condition.
	ContainerNameAndBaselineStatus_UNLOCKED      ContainerNameAndBaselineStatus_BaselineStatus = 2
	ContainerNameAndBaselineStatus_LOCKED        ContainerNameAndBaselineStatus_BaselineStatus = 3
)

// Enum value maps for ContainerNameAndBaselineStatus_BaselineStatus.
var (
	ContainerNameAndBaselineStatus_BaselineStatus_name = map[int32]string{
		0: "INVALID",
		1: "NOT_GENERATED",
		2: "UNLOCKED",
		3: "LOCKED",
	}
	ContainerNameAndBaselineStatus_BaselineStatus_value = map[string]int32{
		"INVALID":       0,
		"NOT_GENERATED": 1,
		"UNLOCKED":      2,
		"LOCKED":        3,
	}
)

func (x ContainerNameAndBaselineStatus_BaselineStatus) Enum() *ContainerNameAndBaselineStatus_BaselineStatus {
	p := new(ContainerNameAndBaselineStatus_BaselineStatus)
	*p = x
	return p
}

func (x ContainerNameAndBaselineStatus_BaselineStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerNameAndBaselineStatus_BaselineStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_process_baseline_proto_enumTypes[0].Descriptor()
}

func (ContainerNameAndBaselineStatus_BaselineStatus) Type() protoreflect.EnumType {
	return &file_storage_process_baseline_proto_enumTypes[0]
}

func (x ContainerNameAndBaselineStatus_BaselineStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerNameAndBaselineStatus_BaselineStatus.Descriptor instead.
func (ContainerNameAndBaselineStatus_BaselineStatus) EnumDescriptor() ([]byte, []int) {
	return file_storage_process_baseline_proto_rawDescGZIP(), []int{4, 0}
}

type ProcessBaselineKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The idea is for the keys to be flexible.
	// Only certain combinations of these will be supported.
	DeploymentId  string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty" search:"Deployment ID,hidden" sql:"type(uuid),index=hash"` // @gotags: search:"Deployment ID,hidden" sql:"type(uuid),index=hash"
	ContainerName string `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	ClusterId     string `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,hidden,store" sql:"type(uuid)"` // @gotags: search:"Cluster ID,hidden,store" sql:"type(uuid)"
	Namespace     string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,hidden,store"`                  // @gotags: search:"Namespace,hidden,store"
}

func (x *ProcessBaselineKey) Reset() {
	*x = ProcessBaselineKey{}
	mi := &file_storage_process_baseline_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessBaselineKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessBaselineKey) ProtoMessage() {}

func (x *ProcessBaselineKey) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_baseline_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessBaselineKey.ProtoReflect.Descriptor instead.
func (*ProcessBaselineKey) Descriptor() ([]byte, []int) {
	return file_storage_process_baseline_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessBaselineKey) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ProcessBaselineKey) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ProcessBaselineKey) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ProcessBaselineKey) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ProcessBaseline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Key                     *ProcessBaselineKey    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Elements                []*BaselineElement     `protobuf:"bytes,3,rep,name=elements,proto3" json:"elements,omitempty"`
	ElementGraveyard        []*BaselineElement     `protobuf:"bytes,8,rep,name=element_graveyard,json=elementGraveyard,proto3" json:"element_graveyard,omitempty" search:"-"` // @gotags: search:"-"
	Created                 *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	UserLockedTimestamp     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=user_locked_timestamp,json=userLockedTimestamp,proto3" json:"user_locked_timestamp,omitempty"`
	StackRoxLockedTimestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=stack_rox_locked_timestamp,json=stackRoxLockedTimestamp,proto3" json:"stack_rox_locked_timestamp,omitempty"`
	LastUpdate              *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

func (x *ProcessBaseline) Reset() {
	*x = ProcessBaseline{}
	mi := &file_storage_process_baseline_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessBaseline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessBaseline) ProtoMessage() {}

func (x *ProcessBaseline) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_baseline_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessBaseline.ProtoReflect.Descriptor instead.
func (*ProcessBaseline) Descriptor() ([]byte, []int) {
	return file_storage_process_baseline_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessBaseline) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProcessBaseline) GetKey() *ProcessBaselineKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ProcessBaseline) GetElements() []*BaselineElement {
	if x != nil {
		return x.Elements
	}
	return nil
}

func (x *ProcessBaseline) GetElementGraveyard() []*BaselineElement {
	if x != nil {
		return x.ElementGraveyard
	}
	return nil
}

func (x *ProcessBaseline) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ProcessBaseline) GetUserLockedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.UserLockedTimestamp
	}
	return nil
}

func (x *ProcessBaseline) GetStackRoxLockedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.StackRoxLockedTimestamp
	}
	return nil
}

func (x *ProcessBaseline) GetLastUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdate
	}
	return nil
}

type BaselineElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Element *BaselineItem `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"`
	Auto    bool          `protobuf:"varint,2,opt,name=auto,proto3" json:"auto,omitempty"`
}

func (x *BaselineElement) Reset() {
	*x = BaselineElement{}
	mi := &file_storage_process_baseline_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaselineElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaselineElement) ProtoMessage() {}

func (x *BaselineElement) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_baseline_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaselineElement.ProtoReflect.Descriptor instead.
func (*BaselineElement) Descriptor() ([]byte, []int) {
	return file_storage_process_baseline_proto_rawDescGZIP(), []int{2}
}

func (x *BaselineElement) GetElement() *BaselineItem {
	if x != nil {
		return x.Element
	}
	return nil
}

func (x *BaselineElement) GetAuto() bool {
	if x != nil {
		return x.Auto
	}
	return false
}

type BaselineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Item:
	//
	//	*BaselineItem_ProcessName
	Item isBaselineItem_Item `protobuf_oneof:"item"`
}

func (x *BaselineItem) Reset() {
	*x = BaselineItem{}
	mi := &file_storage_process_baseline_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaselineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaselineItem) ProtoMessage() {}

func (x *BaselineItem) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_baseline_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaselineItem.ProtoReflect.Descriptor instead.
func (*BaselineItem) Descriptor() ([]byte, []int) {
	return file_storage_process_baseline_proto_rawDescGZIP(), []int{3}
}

func (m *BaselineItem) GetItem() isBaselineItem_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (x *BaselineItem) GetProcessName() string {
	if x, ok := x.GetItem().(*BaselineItem_ProcessName); ok {
		return x.ProcessName
	}
	return ""
}

type isBaselineItem_Item interface {
	isBaselineItem_Item()
}

type BaselineItem_ProcessName struct {
	ProcessName string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3,oneof"`
}

func (*BaselineItem_ProcessName) isBaselineItem_Item() {}

// `ContainerNameAndBaselineStatus` represents a cached result
// of process evaluation on a specific container name.
type ContainerNameAndBaselineStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerName              string                                        `protobuf:"bytes,1,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	BaselineStatus             ContainerNameAndBaselineStatus_BaselineStatus `protobuf:"varint,2,opt,name=baseline_status,json=baselineStatus,proto3,enum=storage.ContainerNameAndBaselineStatus_BaselineStatus" json:"baseline_status,omitempty"`
	AnomalousProcessesExecuted bool                                          `protobuf:"varint,3,opt,name=anomalous_processes_executed,json=anomalousProcessesExecuted,proto3" json:"anomalous_processes_executed,omitempty"`
}

func (x *ContainerNameAndBaselineStatus) Reset() {
	*x = ContainerNameAndBaselineStatus{}
	mi := &file_storage_process_baseline_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerNameAndBaselineStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerNameAndBaselineStatus) ProtoMessage() {}

func (x *ContainerNameAndBaselineStatus) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_baseline_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerNameAndBaselineStatus.ProtoReflect.Descriptor instead.
func (*ContainerNameAndBaselineStatus) Descriptor() ([]byte, []int) {
	return file_storage_process_baseline_proto_rawDescGZIP(), []int{4}
}

func (x *ContainerNameAndBaselineStatus) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ContainerNameAndBaselineStatus) GetBaselineStatus() ContainerNameAndBaselineStatus_BaselineStatus {
	if x != nil {
		return x.BaselineStatus
	}
	return ContainerNameAndBaselineStatus_INVALID
}

func (x *ContainerNameAndBaselineStatus) GetAnomalousProcessesExecuted() bool {
	if x != nil {
		return x.AnomalousProcessesExecuted
	}
	return false
}

// `ProcessBaselineResults` represent cached results of process baseline evaluation.
type ProcessBaselineResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId     string                            `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty" sql:"pk,type(uuid)"` // @gotags: sql:"pk,type(uuid)"
	ClusterId        string                            `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,hidden,store" sql:"type(uuid)"`          // @gotags: search:"Cluster ID,hidden,store"  sql:"type(uuid)"
	Namespace        string                            `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,hidden,store"`                           // @gotags: search:"Namespace,hidden,store"
	BaselineStatuses []*ContainerNameAndBaselineStatus `protobuf:"bytes,2,rep,name=baseline_statuses,json=baselineStatuses,proto3" json:"baseline_statuses,omitempty"`
}

func (x *ProcessBaselineResults) Reset() {
	*x = ProcessBaselineResults{}
	mi := &file_storage_process_baseline_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessBaselineResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessBaselineResults) ProtoMessage() {}

func (x *ProcessBaselineResults) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_baseline_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessBaselineResults.ProtoReflect.Descriptor instead.
func (*ProcessBaselineResults) Descriptor() ([]byte, []int) {
	return file_storage_process_baseline_proto_rawDescGZIP(), []int{5}
}

func (x *ProcessBaselineResults) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ProcessBaselineResults) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ProcessBaselineResults) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ProcessBaselineResults) GetBaselineStatuses() []*ContainerNameAndBaselineStatus {
	if x != nil {
		return x.BaselineStatuses
	}
	return nil
}

var File_storage_process_baseline_proto protoreflect.FileDescriptor

var file_storage_process_baseline_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xe9, 0x03, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a,
	0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67,
	0x72, 0x61, 0x76, 0x65, 0x79, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x47, 0x72, 0x61, 0x76, 0x65, 0x79, 0x61, 0x72, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x4e, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x75, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x57, 0x0a, 0x1a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x78, 0x5f, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x17, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x6f, 0x78, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x56, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x61, 0x75, 0x74, 0x6f, 0x22, 0x3b,
	0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x23,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0xb6, 0x02, 0x0a, 0x1e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64,
	0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5f, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x1c, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c,
	0x6f, 0x75, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x5f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x61, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x47,
	0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e,
	0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b,
	0x45, 0x44, 0x10, 0x03, 0x22, 0xd0, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x54, 0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x42, 0x31, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_storage_process_baseline_proto_rawDescOnce sync.Once
	file_storage_process_baseline_proto_rawDescData = file_storage_process_baseline_proto_rawDesc
)

func file_storage_process_baseline_proto_rawDescGZIP() []byte {
	file_storage_process_baseline_proto_rawDescOnce.Do(func() {
		file_storage_process_baseline_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_process_baseline_proto_rawDescData)
	})
	return file_storage_process_baseline_proto_rawDescData
}

var file_storage_process_baseline_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_process_baseline_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_storage_process_baseline_proto_goTypes = []any{
	(ContainerNameAndBaselineStatus_BaselineStatus)(0), // 0: storage.ContainerNameAndBaselineStatus.BaselineStatus
	(*ProcessBaselineKey)(nil),                         // 1: storage.ProcessBaselineKey
	(*ProcessBaseline)(nil),                            // 2: storage.ProcessBaseline
	(*BaselineElement)(nil),                            // 3: storage.BaselineElement
	(*BaselineItem)(nil),                               // 4: storage.BaselineItem
	(*ContainerNameAndBaselineStatus)(nil),             // 5: storage.ContainerNameAndBaselineStatus
	(*ProcessBaselineResults)(nil),                     // 6: storage.ProcessBaselineResults
	(*timestamppb.Timestamp)(nil),                      // 7: google.protobuf.Timestamp
}
var file_storage_process_baseline_proto_depIdxs = []int32{
	1,  // 0: storage.ProcessBaseline.key:type_name -> storage.ProcessBaselineKey
	3,  // 1: storage.ProcessBaseline.elements:type_name -> storage.BaselineElement
	3,  // 2: storage.ProcessBaseline.element_graveyard:type_name -> storage.BaselineElement
	7,  // 3: storage.ProcessBaseline.created:type_name -> google.protobuf.Timestamp
	7,  // 4: storage.ProcessBaseline.user_locked_timestamp:type_name -> google.protobuf.Timestamp
	7,  // 5: storage.ProcessBaseline.stack_rox_locked_timestamp:type_name -> google.protobuf.Timestamp
	7,  // 6: storage.ProcessBaseline.last_update:type_name -> google.protobuf.Timestamp
	4,  // 7: storage.BaselineElement.element:type_name -> storage.BaselineItem
	0,  // 8: storage.ContainerNameAndBaselineStatus.baseline_status:type_name -> storage.ContainerNameAndBaselineStatus.BaselineStatus
	5,  // 9: storage.ProcessBaselineResults.baseline_statuses:type_name -> storage.ContainerNameAndBaselineStatus
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_storage_process_baseline_proto_init() }
func file_storage_process_baseline_proto_init() {
	if File_storage_process_baseline_proto != nil {
		return
	}
	file_storage_process_baseline_proto_msgTypes[3].OneofWrappers = []any{
		(*BaselineItem_ProcessName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_process_baseline_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_process_baseline_proto_goTypes,
		DependencyIndexes: file_storage_process_baseline_proto_depIdxs,
		EnumInfos:         file_storage_process_baseline_proto_enumTypes,
		MessageInfos:      file_storage_process_baseline_proto_msgTypes,
	}.Build()
	File_storage_process_baseline_proto = out.File
	file_storage_process_baseline_proto_rawDesc = nil
	file_storage_process_baseline_proto_goTypes = nil
	file_storage_process_baseline_proto_depIdxs = nil
}
