// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: storage/report_snapshot.proto

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

type ReportSnapshot_ReportType int32

const (
	ReportSnapshot_VULNERABILITY ReportSnapshot_ReportType = 0
)

// Enum value maps for ReportSnapshot_ReportType.
var (
	ReportSnapshot_ReportType_name = map[int32]string{
		0: "VULNERABILITY",
	}
	ReportSnapshot_ReportType_value = map[string]int32{
		"VULNERABILITY": 0,
	}
)

func (x ReportSnapshot_ReportType) Enum() *ReportSnapshot_ReportType {
	p := new(ReportSnapshot_ReportType)
	*p = x
	return p
}

func (x ReportSnapshot_ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportSnapshot_ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_snapshot_proto_enumTypes[0].Descriptor()
}

func (ReportSnapshot_ReportType) Type() protoreflect.EnumType {
	return &file_storage_report_snapshot_proto_enumTypes[0]
}

func (x ReportSnapshot_ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportSnapshot_ReportType.Descriptor instead.
func (ReportSnapshot_ReportType) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{0, 0}
}

type ReportStatus_RunState int32

const (
	ReportStatus_WAITING   ReportStatus_RunState = 0
	ReportStatus_PREPARING ReportStatus_RunState = 1
	ReportStatus_GENERATED ReportStatus_RunState = 2
	ReportStatus_DELIVERED ReportStatus_RunState = 3
	ReportStatus_FAILURE   ReportStatus_RunState = 4
)

// Enum value maps for ReportStatus_RunState.
var (
	ReportStatus_RunState_name = map[int32]string{
		0: "WAITING",
		1: "PREPARING",
		2: "GENERATED",
		3: "DELIVERED",
		4: "FAILURE",
	}
	ReportStatus_RunState_value = map[string]int32{
		"WAITING":   0,
		"PREPARING": 1,
		"GENERATED": 2,
		"DELIVERED": 3,
		"FAILURE":   4,
	}
)

func (x ReportStatus_RunState) Enum() *ReportStatus_RunState {
	p := new(ReportStatus_RunState)
	*p = x
	return p
}

func (x ReportStatus_RunState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportStatus_RunState) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_snapshot_proto_enumTypes[1].Descriptor()
}

func (ReportStatus_RunState) Type() protoreflect.EnumType {
	return &file_storage_report_snapshot_proto_enumTypes[1]
}

func (x ReportStatus_RunState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportStatus_RunState.Descriptor instead.
func (ReportStatus_RunState) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{3, 0}
}

type ReportStatus_NotificationMethod int32

const (
	ReportStatus_EMAIL    ReportStatus_NotificationMethod = 0
	ReportStatus_DOWNLOAD ReportStatus_NotificationMethod = 1
)

// Enum value maps for ReportStatus_NotificationMethod.
var (
	ReportStatus_NotificationMethod_name = map[int32]string{
		0: "EMAIL",
		1: "DOWNLOAD",
	}
	ReportStatus_NotificationMethod_value = map[string]int32{
		"EMAIL":    0,
		"DOWNLOAD": 1,
	}
)

func (x ReportStatus_NotificationMethod) Enum() *ReportStatus_NotificationMethod {
	p := new(ReportStatus_NotificationMethod)
	*p = x
	return p
}

func (x ReportStatus_NotificationMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportStatus_NotificationMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_snapshot_proto_enumTypes[2].Descriptor()
}

func (ReportStatus_NotificationMethod) Type() protoreflect.EnumType {
	return &file_storage_report_snapshot_proto_enumTypes[2]
}

func (x ReportStatus_NotificationMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportStatus_NotificationMethod.Descriptor instead.
func (ReportStatus_NotificationMethod) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{3, 1}
}

type ReportStatus_RunMethod int32

const (
	ReportStatus_ON_DEMAND ReportStatus_RunMethod = 0
	ReportStatus_SCHEDULED ReportStatus_RunMethod = 1
)

// Enum value maps for ReportStatus_RunMethod.
var (
	ReportStatus_RunMethod_name = map[int32]string{
		0: "ON_DEMAND",
		1: "SCHEDULED",
	}
	ReportStatus_RunMethod_value = map[string]int32{
		"ON_DEMAND": 0,
		"SCHEDULED": 1,
	}
)

func (x ReportStatus_RunMethod) Enum() *ReportStatus_RunMethod {
	p := new(ReportStatus_RunMethod)
	*p = x
	return p
}

func (x ReportStatus_RunMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportStatus_RunMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_snapshot_proto_enumTypes[3].Descriptor()
}

func (ReportStatus_RunMethod) Type() protoreflect.EnumType {
	return &file_storage_report_snapshot_proto_enumTypes[3]
}

func (x ReportStatus_RunMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportStatus_RunMethod.Descriptor instead.
func (ReportStatus_RunMethod) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{3, 2}
}

// ReportSnapshot stores the snapshot of a report job. It stores a projection of ReportConfiguration, collection,
// vulnerability filters, notifiers, etc used to generate a report. It also stores the final status of the report job.
type ReportSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId              string                    `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty" sql:"pk,type(uuid)"`                                          // @gotags: sql:"pk,type(uuid)"
	ReportConfigurationId string                    `protobuf:"bytes,2,opt,name=report_configuration_id,json=reportConfigurationId,proto3" json:"report_configuration_id,omitempty" search:"Report Configuration ID" sql:"fk(ReportConfiguration:id)"` // @gotags: search:"Report Configuration ID" sql:"fk(ReportConfiguration:id)"
	Name                  string                    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" search:"Report Name"`                                                                  // @gotags: search:"Report Name"
	Description           string                    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Type                  ReportSnapshot_ReportType `protobuf:"varint,5,opt,name=type,proto3,enum=storage.ReportSnapshot_ReportType" json:"type,omitempty"`
	// Types that are assignable to Filter:
	//
	//	*ReportSnapshot_VulnReportFilters
	Filter       isReportSnapshot_Filter `protobuf_oneof:"filter"`
	Collection   *CollectionSnapshot     `protobuf:"bytes,7,opt,name=collection,proto3" json:"collection,omitempty"`
	Schedule     *Schedule               `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
	ReportStatus *ReportStatus           `protobuf:"bytes,9,opt,name=report_status,json=reportStatus,proto3" json:"report_status,omitempty"`
	Notifiers    []*NotifierSnapshot     `protobuf:"bytes,10,rep,name=notifiers,proto3" json:"notifiers,omitempty"`
	Requester    *SlimUser               `protobuf:"bytes,11,opt,name=requester,proto3" json:"requester,omitempty"`
}

func (x *ReportSnapshot) Reset() {
	*x = ReportSnapshot{}
	mi := &file_storage_report_snapshot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSnapshot) ProtoMessage() {}

func (x *ReportSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_snapshot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSnapshot.ProtoReflect.Descriptor instead.
func (*ReportSnapshot) Descriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *ReportSnapshot) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *ReportSnapshot) GetReportConfigurationId() string {
	if x != nil {
		return x.ReportConfigurationId
	}
	return ""
}

func (x *ReportSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportSnapshot) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReportSnapshot) GetType() ReportSnapshot_ReportType {
	if x != nil {
		return x.Type
	}
	return ReportSnapshot_VULNERABILITY
}

func (m *ReportSnapshot) GetFilter() isReportSnapshot_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (x *ReportSnapshot) GetVulnReportFilters() *VulnerabilityReportFilters {
	if x, ok := x.GetFilter().(*ReportSnapshot_VulnReportFilters); ok {
		return x.VulnReportFilters
	}
	return nil
}

func (x *ReportSnapshot) GetCollection() *CollectionSnapshot {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *ReportSnapshot) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *ReportSnapshot) GetReportStatus() *ReportStatus {
	if x != nil {
		return x.ReportStatus
	}
	return nil
}

func (x *ReportSnapshot) GetNotifiers() []*NotifierSnapshot {
	if x != nil {
		return x.Notifiers
	}
	return nil
}

func (x *ReportSnapshot) GetRequester() *SlimUser {
	if x != nil {
		return x.Requester
	}
	return nil
}

type isReportSnapshot_Filter interface {
	isReportSnapshot_Filter()
}

type ReportSnapshot_VulnReportFilters struct {
	VulnReportFilters *VulnerabilityReportFilters `protobuf:"bytes,6,opt,name=vuln_report_filters,json=vulnReportFilters,proto3,oneof"`
}

func (*ReportSnapshot_VulnReportFilters) isReportSnapshot_Filter() {}

type CollectionSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CollectionSnapshot) Reset() {
	*x = CollectionSnapshot{}
	mi := &file_storage_report_snapshot_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectionSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionSnapshot) ProtoMessage() {}

func (x *CollectionSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_snapshot_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionSnapshot.ProtoReflect.Descriptor instead.
func (*CollectionSnapshot) Descriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CollectionSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NotifierSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NotifierConfig:
	//
	//	*NotifierSnapshot_EmailConfig
	NotifierConfig isNotifierSnapshot_NotifierConfig `protobuf_oneof:"notifier_config"`
	NotifierName   string                            `protobuf:"bytes,2,opt,name=notifier_name,json=notifierName,proto3" json:"notifier_name,omitempty"`
}

func (x *NotifierSnapshot) Reset() {
	*x = NotifierSnapshot{}
	mi := &file_storage_report_snapshot_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifierSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifierSnapshot) ProtoMessage() {}

func (x *NotifierSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_snapshot_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifierSnapshot.ProtoReflect.Descriptor instead.
func (*NotifierSnapshot) Descriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{2}
}

func (m *NotifierSnapshot) GetNotifierConfig() isNotifierSnapshot_NotifierConfig {
	if m != nil {
		return m.NotifierConfig
	}
	return nil
}

func (x *NotifierSnapshot) GetEmailConfig() *EmailNotifierConfiguration {
	if x, ok := x.GetNotifierConfig().(*NotifierSnapshot_EmailConfig); ok {
		return x.EmailConfig
	}
	return nil
}

func (x *NotifierSnapshot) GetNotifierName() string {
	if x != nil {
		return x.NotifierName
	}
	return ""
}

type isNotifierSnapshot_NotifierConfig interface {
	isNotifierSnapshot_NotifierConfig()
}

type NotifierSnapshot_EmailConfig struct {
	EmailConfig *EmailNotifierConfiguration `protobuf:"bytes,1,opt,name=email_config,json=emailConfig,proto3,oneof"`
}

func (*NotifierSnapshot_EmailConfig) isNotifierSnapshot_NotifierConfig() {}

type ReportStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunState                 ReportStatus_RunState           `protobuf:"varint,1,opt,name=run_state,json=runState,proto3,enum=storage.ReportStatus_RunState" json:"run_state,omitempty" search:"Report State"` // @gotags: search:"Report State"
	QueuedAt                 *timestamppb.Timestamp          `protobuf:"bytes,2,opt,name=queued_at,json=queuedAt,proto3" json:"queued_at,omitempty" search:"Report Init Time"`                                     // @gotags: search:"Report Init Time"
	CompletedAt              *timestamppb.Timestamp          `protobuf:"bytes,3,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty" search:"Report Completion Time"`                            // @gotags: search:"Report Completion Time"
	ErrorMsg                 string                          `protobuf:"bytes,4,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	ReportRequestType        ReportStatus_RunMethod          `protobuf:"varint,5,opt,name=report_request_type,json=reportRequestType,proto3,enum=storage.ReportStatus_RunMethod" json:"report_request_type,omitempty" search:"Report Request Type"`                               // @gotags: search:"Report Request Type"
	ReportNotificationMethod ReportStatus_NotificationMethod `protobuf:"varint,6,opt,name=report_notification_method,json=reportNotificationMethod,proto3,enum=storage.ReportStatus_NotificationMethod" json:"report_notification_method,omitempty" search:"Report Notification Method"` // @gotags: search:"Report Notification Method"
}

func (x *ReportStatus) Reset() {
	*x = ReportStatus{}
	mi := &file_storage_report_snapshot_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportStatus) ProtoMessage() {}

func (x *ReportStatus) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_snapshot_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportStatus.ProtoReflect.Descriptor instead.
func (*ReportStatus) Descriptor() ([]byte, []int) {
	return file_storage_report_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *ReportStatus) GetRunState() ReportStatus_RunState {
	if x != nil {
		return x.RunState
	}
	return ReportStatus_WAITING
}

func (x *ReportStatus) GetQueuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.QueuedAt
	}
	return nil
}

func (x *ReportStatus) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *ReportStatus) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *ReportStatus) GetReportRequestType() ReportStatus_RunMethod {
	if x != nil {
		return x.ReportRequestType
	}
	return ReportStatus_ON_DEMAND
}

func (x *ReportStatus) GetReportNotificationMethod() ReportStatus_NotificationMethod {
	if x != nil {
		return x.ReportNotificationMethod
	}
	return ReportStatus_EMAIL
}

var File_storage_report_snapshot_proto protoreflect.FileDescriptor

var file_storage_report_snapshot_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x04, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x13,
	0x76, 0x75, 0x6c, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00,
	0x52, 0x11, 0x76, 0x75, 0x6c, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x6c, 0x69, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x55, 0x4c, 0x4e, 0x45, 0x52, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x10, 0x00, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x38, 0x0a, 0x12, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x10, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12,
	0x48, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x11,
	0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0xc6, 0x04, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x37, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x4f, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x66, 0x0a, 0x1a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x18, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x51, 0x0a,
	0x08, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x04,
	0x22, 0x2d, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01, 0x22,
	0x29, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0d, 0x0a, 0x09,
	0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_storage_report_snapshot_proto_rawDescOnce sync.Once
	file_storage_report_snapshot_proto_rawDescData = file_storage_report_snapshot_proto_rawDesc
)

func file_storage_report_snapshot_proto_rawDescGZIP() []byte {
	file_storage_report_snapshot_proto_rawDescOnce.Do(func() {
		file_storage_report_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_report_snapshot_proto_rawDescData)
	})
	return file_storage_report_snapshot_proto_rawDescData
}

var file_storage_report_snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_storage_report_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_report_snapshot_proto_goTypes = []any{
	(ReportSnapshot_ReportType)(0),       // 0: storage.ReportSnapshot.ReportType
	(ReportStatus_RunState)(0),           // 1: storage.ReportStatus.RunState
	(ReportStatus_NotificationMethod)(0), // 2: storage.ReportStatus.NotificationMethod
	(ReportStatus_RunMethod)(0),          // 3: storage.ReportStatus.RunMethod
	(*ReportSnapshot)(nil),               // 4: storage.ReportSnapshot
	(*CollectionSnapshot)(nil),           // 5: storage.CollectionSnapshot
	(*NotifierSnapshot)(nil),             // 6: storage.NotifierSnapshot
	(*ReportStatus)(nil),                 // 7: storage.ReportStatus
	(*VulnerabilityReportFilters)(nil),   // 8: storage.VulnerabilityReportFilters
	(*Schedule)(nil),                     // 9: storage.Schedule
	(*SlimUser)(nil),                     // 10: storage.SlimUser
	(*EmailNotifierConfiguration)(nil),   // 11: storage.EmailNotifierConfiguration
	(*timestamppb.Timestamp)(nil),        // 12: google.protobuf.Timestamp
}
var file_storage_report_snapshot_proto_depIdxs = []int32{
	0,  // 0: storage.ReportSnapshot.type:type_name -> storage.ReportSnapshot.ReportType
	8,  // 1: storage.ReportSnapshot.vuln_report_filters:type_name -> storage.VulnerabilityReportFilters
	5,  // 2: storage.ReportSnapshot.collection:type_name -> storage.CollectionSnapshot
	9,  // 3: storage.ReportSnapshot.schedule:type_name -> storage.Schedule
	7,  // 4: storage.ReportSnapshot.report_status:type_name -> storage.ReportStatus
	6,  // 5: storage.ReportSnapshot.notifiers:type_name -> storage.NotifierSnapshot
	10, // 6: storage.ReportSnapshot.requester:type_name -> storage.SlimUser
	11, // 7: storage.NotifierSnapshot.email_config:type_name -> storage.EmailNotifierConfiguration
	1,  // 8: storage.ReportStatus.run_state:type_name -> storage.ReportStatus.RunState
	12, // 9: storage.ReportStatus.queued_at:type_name -> google.protobuf.Timestamp
	12, // 10: storage.ReportStatus.completed_at:type_name -> google.protobuf.Timestamp
	3,  // 11: storage.ReportStatus.report_request_type:type_name -> storage.ReportStatus.RunMethod
	2,  // 12: storage.ReportStatus.report_notification_method:type_name -> storage.ReportStatus.NotificationMethod
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_storage_report_snapshot_proto_init() }
func file_storage_report_snapshot_proto_init() {
	if File_storage_report_snapshot_proto != nil {
		return
	}
	file_storage_report_configuration_proto_init()
	file_storage_report_notifier_configuration_proto_init()
	file_storage_schedule_proto_init()
	file_storage_user_proto_init()
	file_storage_report_snapshot_proto_msgTypes[0].OneofWrappers = []any{
		(*ReportSnapshot_VulnReportFilters)(nil),
	}
	file_storage_report_snapshot_proto_msgTypes[2].OneofWrappers = []any{
		(*NotifierSnapshot_EmailConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_report_snapshot_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_report_snapshot_proto_goTypes,
		DependencyIndexes: file_storage_report_snapshot_proto_depIdxs,
		EnumInfos:         file_storage_report_snapshot_proto_enumTypes,
		MessageInfos:      file_storage_report_snapshot_proto_msgTypes,
	}.Build()
	File_storage_report_snapshot_proto = out.File
	file_storage_report_snapshot_proto_rawDesc = nil
	file_storage_report_snapshot_proto_goTypes = nil
	file_storage_report_snapshot_proto_depIdxs = nil
}
