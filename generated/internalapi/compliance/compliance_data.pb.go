// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: internalapi/compliance/compliance_data.proto

package compliance

import (
	storage "github.com/stackrox/rox/generated/storage"
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

// Next Available Tag: 2
type GZIPDataChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gzip []byte `protobuf:"bytes,1,opt,name=gzip,proto3" json:"gzip,omitempty"`
}

func (x *GZIPDataChunk) Reset() {
	*x = GZIPDataChunk{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GZIPDataChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GZIPDataChunk) ProtoMessage() {}

func (x *GZIPDataChunk) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GZIPDataChunk.ProtoReflect.Descriptor instead.
func (*GZIPDataChunk) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{0}
}

func (x *GZIPDataChunk) GetGzip() []byte {
	if x != nil {
		return x.Gzip
	}
	return nil
}

// Next Available Tag: 8
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	User        uint32  `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	UserName    string  `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Group       uint32  `protobuf:"varint,4,opt,name=group,proto3" json:"group,omitempty"`
	GroupName   string  `protobuf:"bytes,5,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Permissions uint32  `protobuf:"varint,6,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Content     []byte  `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	IsDir       bool    `protobuf:"varint,8,opt,name=is_dir,json=isDir,proto3" json:"is_dir,omitempty"`
	Children    []*File `protobuf:"bytes,9,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetUser() uint32 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *File) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *File) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *File) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *File) GetPermissions() uint32 {
	if x != nil {
		return x.Permissions
	}
	return 0
}

func (x *File) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *File) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

func (x *File) GetChildren() []*File {
	if x != nil {
		return x.Children
	}
	return nil
}

// Next Available Tag: 3
type CommandLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Process string              `protobuf:"bytes,1,opt,name=process,proto3" json:"process,omitempty"`
	Args    []*CommandLine_Args `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *CommandLine) Reset() {
	*x = CommandLine{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandLine) ProtoMessage() {}

func (x *CommandLine) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandLine.ProtoReflect.Descriptor instead.
func (*CommandLine) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{2}
}

func (x *CommandLine) GetProcess() string {
	if x != nil {
		return x.Process
	}
	return ""
}

func (x *CommandLine) GetArgs() []*CommandLine_Args {
	if x != nil {
		return x.Args
	}
	return nil
}

type InsecureRegistriesConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InsecureRegistries []string `protobuf:"bytes,1,rep,name=insecure_registries,json=insecureRegistries,proto3" json:"insecure_registries,omitempty"`
	InsecureCidrs      []string `protobuf:"bytes,2,rep,name=insecure_cidrs,json=insecureCidrs,proto3" json:"insecure_cidrs,omitempty"`
}

func (x *InsecureRegistriesConfig) Reset() {
	*x = InsecureRegistriesConfig{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsecureRegistriesConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsecureRegistriesConfig) ProtoMessage() {}

func (x *InsecureRegistriesConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsecureRegistriesConfig.ProtoReflect.Descriptor instead.
func (*InsecureRegistriesConfig) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{3}
}

func (x *InsecureRegistriesConfig) GetInsecureRegistries() []string {
	if x != nil {
		return x.InsecureRegistries
	}
	return nil
}

func (x *InsecureRegistriesConfig) GetInsecureCidrs() []string {
	if x != nil {
		return x.InsecureCidrs
	}
	return nil
}

type ContainerRuntimeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InsecureRegistries *InsecureRegistriesConfig `protobuf:"bytes,1,opt,name=insecure_registries,json=insecureRegistries,proto3" json:"insecure_registries,omitempty"`
}

func (x *ContainerRuntimeInfo) Reset() {
	*x = ContainerRuntimeInfo{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerRuntimeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerRuntimeInfo) ProtoMessage() {}

func (x *ContainerRuntimeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerRuntimeInfo.ProtoReflect.Descriptor instead.
func (*ContainerRuntimeInfo) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{4}
}

func (x *ContainerRuntimeInfo) GetInsecureRegistries() *InsecureRegistriesConfig {
	if x != nil {
		return x.InsecureRegistries
	}
	return nil
}

type ComplianceStandardResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeCheckResults    map[string]*storage.ComplianceResultValue `protobuf:"bytes,1,rep,name=node_check_results,json=nodeCheckResults,proto3" json:"node_check_results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClusterCheckResults map[string]*storage.ComplianceResultValue `protobuf:"bytes,2,rep,name=cluster_check_results,json=clusterCheckResults,proto3" json:"cluster_check_results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ComplianceStandardResult) Reset() {
	*x = ComplianceStandardResult{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceStandardResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceStandardResult) ProtoMessage() {}

func (x *ComplianceStandardResult) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceStandardResult.ProtoReflect.Descriptor instead.
func (*ComplianceStandardResult) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{5}
}

func (x *ComplianceStandardResult) GetNodeCheckResults() map[string]*storage.ComplianceResultValue {
	if x != nil {
		return x.NodeCheckResults
	}
	return nil
}

func (x *ComplianceStandardResult) GetClusterCheckResults() map[string]*storage.ComplianceResultValue {
	if x != nil {
		return x.ClusterCheckResults
	}
	return nil
}

// Next Available Tag: 11
type ComplianceReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName             string                  `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ScrapeId             string                  `protobuf:"bytes,2,opt,name=scrape_id,json=scrapeId,proto3" json:"scrape_id,omitempty"`
	DockerData           *GZIPDataChunk          `protobuf:"bytes,3,opt,name=docker_data,json=dockerData,proto3" json:"docker_data,omitempty"`
	CommandLines         map[string]*CommandLine `protobuf:"bytes,4,rep,name=command_lines,json=commandLines,proto3" json:"command_lines,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Files                map[string]*File        `protobuf:"bytes,5,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SystemdFiles         map[string]*File        `protobuf:"bytes,6,rep,name=systemd_files,json=systemdFiles,proto3" json:"systemd_files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ContainerRuntimeInfo *ContainerRuntimeInfo   `protobuf:"bytes,9,opt,name=container_runtime_info,json=containerRuntimeInfo,proto3" json:"container_runtime_info,omitempty"`
	Time                 *timestamppb.Timestamp  `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Evidence             *GZIPDataChunk          `protobuf:"bytes,10,opt,name=evidence,proto3" json:"evidence,omitempty"`
}

func (x *ComplianceReturn) Reset() {
	*x = ComplianceReturn{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceReturn) ProtoMessage() {}

func (x *ComplianceReturn) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceReturn.ProtoReflect.Descriptor instead.
func (*ComplianceReturn) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{6}
}

func (x *ComplianceReturn) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *ComplianceReturn) GetScrapeId() string {
	if x != nil {
		return x.ScrapeId
	}
	return ""
}

func (x *ComplianceReturn) GetDockerData() *GZIPDataChunk {
	if x != nil {
		return x.DockerData
	}
	return nil
}

func (x *ComplianceReturn) GetCommandLines() map[string]*CommandLine {
	if x != nil {
		return x.CommandLines
	}
	return nil
}

func (x *ComplianceReturn) GetFiles() map[string]*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *ComplianceReturn) GetSystemdFiles() map[string]*File {
	if x != nil {
		return x.SystemdFiles
	}
	return nil
}

func (x *ComplianceReturn) GetContainerRuntimeInfo() *ContainerRuntimeInfo {
	if x != nil {
		return x.ContainerRuntimeInfo
	}
	return nil
}

func (x *ComplianceReturn) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ComplianceReturn) GetEvidence() *GZIPDataChunk {
	if x != nil {
		return x.Evidence
	}
	return nil
}

type CommandLine_Args struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	File   *File    `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"` // This will be set if it looks like the value defines a file and we found the file
}

func (x *CommandLine_Args) Reset() {
	*x = CommandLine_Args{}
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandLine_Args) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandLine_Args) ProtoMessage() {}

func (x *CommandLine_Args) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_compliance_compliance_data_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandLine_Args.ProtoReflect.Descriptor instead.
func (*CommandLine_Args) Descriptor() ([]byte, []int) {
	return file_internalapi_compliance_compliance_data_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CommandLine_Args) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CommandLine_Args) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *CommandLine_Args) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

var File_internalapi_compliance_compliance_data_proto protoreflect.FileDescriptor

var file_internalapi_compliance_compliance_data_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x5a, 0x49, 0x50, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x7a, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x67, 0x7a, 0x69, 0x70, 0x22, 0x81, 0x02, 0x0a, 0x04, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x64, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72,
	0x12, 0x2c, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xb1,
	0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x2e,
	0x41, 0x72, 0x67, 0x73, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x1a, 0x56, 0x0a, 0x04, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x72, 0x0a, 0x18, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f,
	0x0a, 0x13, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x43, 0x69, 0x64, 0x72, 0x73, 0x22, 0x6d, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x55,
	0x0a, 0x13, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xc4, 0x03, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x68, 0x0a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x6e, 0x6f, 0x64, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x71, 0x0a, 0x15,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a,
	0x63, 0x0a, 0x15, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x66, 0x0a, 0x18, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa9, 0x06, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x5a,
	0x49, 0x50, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x0a, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x53, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x0d, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x56, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x5a, 0x49, 0x50, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x1a,
	0x58, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4a, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x43, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5a, 0x23, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_compliance_compliance_data_proto_rawDescOnce sync.Once
	file_internalapi_compliance_compliance_data_proto_rawDescData = file_internalapi_compliance_compliance_data_proto_rawDesc
)

func file_internalapi_compliance_compliance_data_proto_rawDescGZIP() []byte {
	file_internalapi_compliance_compliance_data_proto_rawDescOnce.Do(func() {
		file_internalapi_compliance_compliance_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_compliance_compliance_data_proto_rawDescData)
	})
	return file_internalapi_compliance_compliance_data_proto_rawDescData
}

var file_internalapi_compliance_compliance_data_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_internalapi_compliance_compliance_data_proto_goTypes = []any{
	(*GZIPDataChunk)(nil),                 // 0: compliance.GZIPDataChunk
	(*File)(nil),                          // 1: compliance.File
	(*CommandLine)(nil),                   // 2: compliance.CommandLine
	(*InsecureRegistriesConfig)(nil),      // 3: compliance.InsecureRegistriesConfig
	(*ContainerRuntimeInfo)(nil),          // 4: compliance.ContainerRuntimeInfo
	(*ComplianceStandardResult)(nil),      // 5: compliance.ComplianceStandardResult
	(*ComplianceReturn)(nil),              // 6: compliance.ComplianceReturn
	(*CommandLine_Args)(nil),              // 7: compliance.CommandLine.Args
	nil,                                   // 8: compliance.ComplianceStandardResult.NodeCheckResultsEntry
	nil,                                   // 9: compliance.ComplianceStandardResult.ClusterCheckResultsEntry
	nil,                                   // 10: compliance.ComplianceReturn.CommandLinesEntry
	nil,                                   // 11: compliance.ComplianceReturn.FilesEntry
	nil,                                   // 12: compliance.ComplianceReturn.SystemdFilesEntry
	(*timestamppb.Timestamp)(nil),         // 13: google.protobuf.Timestamp
	(*storage.ComplianceResultValue)(nil), // 14: storage.ComplianceResultValue
}
var file_internalapi_compliance_compliance_data_proto_depIdxs = []int32{
	1,  // 0: compliance.File.children:type_name -> compliance.File
	7,  // 1: compliance.CommandLine.args:type_name -> compliance.CommandLine.Args
	3,  // 2: compliance.ContainerRuntimeInfo.insecure_registries:type_name -> compliance.InsecureRegistriesConfig
	8,  // 3: compliance.ComplianceStandardResult.node_check_results:type_name -> compliance.ComplianceStandardResult.NodeCheckResultsEntry
	9,  // 4: compliance.ComplianceStandardResult.cluster_check_results:type_name -> compliance.ComplianceStandardResult.ClusterCheckResultsEntry
	0,  // 5: compliance.ComplianceReturn.docker_data:type_name -> compliance.GZIPDataChunk
	10, // 6: compliance.ComplianceReturn.command_lines:type_name -> compliance.ComplianceReturn.CommandLinesEntry
	11, // 7: compliance.ComplianceReturn.files:type_name -> compliance.ComplianceReturn.FilesEntry
	12, // 8: compliance.ComplianceReturn.systemd_files:type_name -> compliance.ComplianceReturn.SystemdFilesEntry
	4,  // 9: compliance.ComplianceReturn.container_runtime_info:type_name -> compliance.ContainerRuntimeInfo
	13, // 10: compliance.ComplianceReturn.time:type_name -> google.protobuf.Timestamp
	0,  // 11: compliance.ComplianceReturn.evidence:type_name -> compliance.GZIPDataChunk
	1,  // 12: compliance.CommandLine.Args.file:type_name -> compliance.File
	14, // 13: compliance.ComplianceStandardResult.NodeCheckResultsEntry.value:type_name -> storage.ComplianceResultValue
	14, // 14: compliance.ComplianceStandardResult.ClusterCheckResultsEntry.value:type_name -> storage.ComplianceResultValue
	2,  // 15: compliance.ComplianceReturn.CommandLinesEntry.value:type_name -> compliance.CommandLine
	1,  // 16: compliance.ComplianceReturn.FilesEntry.value:type_name -> compliance.File
	1,  // 17: compliance.ComplianceReturn.SystemdFilesEntry.value:type_name -> compliance.File
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_internalapi_compliance_compliance_data_proto_init() }
func file_internalapi_compliance_compliance_data_proto_init() {
	if File_internalapi_compliance_compliance_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_compliance_compliance_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_compliance_compliance_data_proto_goTypes,
		DependencyIndexes: file_internalapi_compliance_compliance_data_proto_depIdxs,
		MessageInfos:      file_internalapi_compliance_compliance_data_proto_msgTypes,
	}.Build()
	File_internalapi_compliance_compliance_data_proto = out.File
	file_internalapi_compliance_compliance_data_proto_rawDesc = nil
	file_internalapi_compliance_compliance_data_proto_goTypes = nil
	file_internalapi_compliance_compliance_data_proto_depIdxs = nil
}
