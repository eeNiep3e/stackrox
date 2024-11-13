// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: storage/role.proto

package storage

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

type Access int32

const (
	Access_NO_ACCESS         Access = 0
	Access_READ_ACCESS       Access = 1
	Access_READ_WRITE_ACCESS Access = 2
)

// Enum value maps for Access.
var (
	Access_name = map[int32]string{
		0: "NO_ACCESS",
		1: "READ_ACCESS",
		2: "READ_WRITE_ACCESS",
	}
	Access_value = map[string]int32{
		"NO_ACCESS":         0,
		"READ_ACCESS":       1,
		"READ_WRITE_ACCESS": 2,
	}
)

func (x Access) Enum() *Access {
	p := new(Access)
	*p = x
	return p
}

func (x Access) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Access) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_role_proto_enumTypes[0].Descriptor()
}

func (Access) Type() protoreflect.EnumType {
	return &file_storage_role_proto_enumTypes[0]
}

func (x Access) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Access.Descriptor instead.
func (Access) EnumDescriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{0}
}

type EffectiveAccessScope_State int32

const (
	EffectiveAccessScope_UNKNOWN  EffectiveAccessScope_State = 0
	EffectiveAccessScope_INCLUDED EffectiveAccessScope_State = 1
	EffectiveAccessScope_EXCLUDED EffectiveAccessScope_State = 2
	EffectiveAccessScope_PARTIAL  EffectiveAccessScope_State = 3
)

// Enum value maps for EffectiveAccessScope_State.
var (
	EffectiveAccessScope_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "INCLUDED",
		2: "EXCLUDED",
		3: "PARTIAL",
	}
	EffectiveAccessScope_State_value = map[string]int32{
		"UNKNOWN":  0,
		"INCLUDED": 1,
		"EXCLUDED": 2,
		"PARTIAL":  3,
	}
)

func (x EffectiveAccessScope_State) Enum() *EffectiveAccessScope_State {
	p := new(EffectiveAccessScope_State)
	*p = x
	return p
}

func (x EffectiveAccessScope_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EffectiveAccessScope_State) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_role_proto_enumTypes[1].Descriptor()
}

func (EffectiveAccessScope_State) Type() protoreflect.EnumType {
	return &file_storage_role_proto_enumTypes[1]
}

func (x EffectiveAccessScope_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EffectiveAccessScope_State.Descriptor instead.
func (EffectiveAccessScope_State) EnumDescriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{3, 0}
}

// A role specifies which actions are allowed for which subset of cluster
// objects. Permissions be can either specified directly via setting
// resource_to_access together with global_access or by referencing a
// permission set by its id in permission_set_name.
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `name` and `description` are provided by the user and can be changed.
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The associated PermissionSet and AccessScope for this Role.
	PermissionSetId string `protobuf:"bytes,6,opt,name=permission_set_id,json=permissionSetId,proto3" json:"permission_set_id,omitempty" sql:"type(uuid)"` // @gotags: sql:"type(uuid)"
	AccessScopeId   string `protobuf:"bytes,7,opt,name=access_scope_id,json=accessScopeId,proto3" json:"access_scope_id,omitempty" sql:"type(uuid)"`       // @gotags: sql:"type(uuid)"
	// Minimum (not default!) access level for every resource. Can be extended
	// below by explicit permissions but not shrunk.
	// Deprecated 2021-04-20 in favor of `permission_set_id`.
	// This field now should be always NO_ACCESS
	//
	// Deprecated: Marked as deprecated in storage/role.proto.
	GlobalAccess Access `protobuf:"varint,2,opt,name=global_access,json=globalAccess,proto3,enum=storage.Access" json:"global_access,omitempty"`
	// Deprecated 2021-04-20 in favor of `permission_set_id`.
	//
	// Deprecated: Marked as deprecated in storage/role.proto.
	ResourceToAccess map[string]Access `protobuf:"bytes,3,rep,name=resource_to_access,json=resourceToAccess,proto3" json:"resource_to_access,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=storage.Access"`
	Traits           *Traits           `protobuf:"bytes,8,opt,name=traits,proto3" json:"traits,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_storage_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetPermissionSetId() string {
	if x != nil {
		return x.PermissionSetId
	}
	return ""
}

func (x *Role) GetAccessScopeId() string {
	if x != nil {
		return x.AccessScopeId
	}
	return ""
}

// Deprecated: Marked as deprecated in storage/role.proto.
func (x *Role) GetGlobalAccess() Access {
	if x != nil {
		return x.GlobalAccess
	}
	return Access_NO_ACCESS
}

// Deprecated: Marked as deprecated in storage/role.proto.
func (x *Role) GetResourceToAccess() map[string]Access {
	if x != nil {
		return x.ResourceToAccess
	}
	return nil
}

func (x *Role) GetTraits() *Traits {
	if x != nil {
		return x.Traits
	}
	return nil
}

// This encodes a set of permissions for StackRox resources.
type PermissionSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is generated and cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"` // @gotags: sql:"pk,type(uuid)"
	// `name` and `description` are provided by the user and can be changed.
	Name             string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sql:"unique"` // @gotags: sql:"unique"
	Description      string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ResourceToAccess map[string]Access `protobuf:"bytes,4,rep,name=resource_to_access,json=resourceToAccess,proto3" json:"resource_to_access,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=storage.Access"`
	Traits           *Traits           `protobuf:"bytes,5,opt,name=traits,proto3" json:"traits,omitempty"`
}

func (x *PermissionSet) Reset() {
	*x = PermissionSet{}
	mi := &file_storage_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionSet) ProtoMessage() {}

func (x *PermissionSet) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionSet.ProtoReflect.Descriptor instead.
func (*PermissionSet) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionSet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PermissionSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionSet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PermissionSet) GetResourceToAccess() map[string]Access {
	if x != nil {
		return x.ResourceToAccess
	}
	return nil
}

func (x *PermissionSet) GetTraits() *Traits {
	if x != nil {
		return x.Traits
	}
	return nil
}

// Simple access scope is a (simple) selection criteria for scoped resources.
// It does *not* allow multi-component AND-rules nor set operations on names.
type SimpleAccessScope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `id` is generated and cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"` // @gotags: sql:"pk,type(uuid)"
	// `name` and `description` are provided by the user and can be changed.
	Name        string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sql:"unique"` // @gotags: sql:"unique"
	Description string                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Rules       *SimpleAccessScope_Rules `protobuf:"bytes,4,opt,name=rules,proto3" json:"rules,omitempty"`
	Traits      *Traits                  `protobuf:"bytes,5,opt,name=traits,proto3" json:"traits,omitempty"`
}

func (x *SimpleAccessScope) Reset() {
	*x = SimpleAccessScope{}
	mi := &file_storage_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleAccessScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleAccessScope) ProtoMessage() {}

func (x *SimpleAccessScope) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleAccessScope.ProtoReflect.Descriptor instead.
func (*SimpleAccessScope) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleAccessScope) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimpleAccessScope) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleAccessScope) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SimpleAccessScope) GetRules() *SimpleAccessScope_Rules {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *SimpleAccessScope) GetTraits() *Traits {
	if x != nil {
		return x.Traits
	}
	return nil
}

// EffectiveAccessScope describes which clusters and namespaces are "in scope"
// given current state. Basically, if AccessScope is applied to the currently
// known clusters and namespaces, the result is EffectiveAccessScope.
//
// EffectiveAccessScope represents a tree with nodes marked as included and
// excluded. If a node is included, all its child nodes are included.
type EffectiveAccessScope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*EffectiveAccessScope_Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *EffectiveAccessScope) Reset() {
	*x = EffectiveAccessScope{}
	mi := &file_storage_role_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EffectiveAccessScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveAccessScope) ProtoMessage() {}

func (x *EffectiveAccessScope) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveAccessScope.ProtoReflect.Descriptor instead.
func (*EffectiveAccessScope) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{3}
}

func (x *EffectiveAccessScope) GetClusters() []*EffectiveAccessScope_Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// Each element of any repeated field is an individual rule. Rules are
// joined by logical OR: if there exists a rule allowing resource `x`,
// `x` is in the access scope.
type SimpleAccessScope_Rules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludedClusters        []string                             `protobuf:"bytes,1,rep,name=included_clusters,json=includedClusters,proto3" json:"included_clusters,omitempty"`
	IncludedNamespaces      []*SimpleAccessScope_Rules_Namespace `protobuf:"bytes,2,rep,name=included_namespaces,json=includedNamespaces,proto3" json:"included_namespaces,omitempty"`
	ClusterLabelSelectors   []*SetBasedLabelSelector             `protobuf:"bytes,3,rep,name=cluster_label_selectors,json=clusterLabelSelectors,proto3" json:"cluster_label_selectors,omitempty"`
	NamespaceLabelSelectors []*SetBasedLabelSelector             `protobuf:"bytes,4,rep,name=namespace_label_selectors,json=namespaceLabelSelectors,proto3" json:"namespace_label_selectors,omitempty"`
}

func (x *SimpleAccessScope_Rules) Reset() {
	*x = SimpleAccessScope_Rules{}
	mi := &file_storage_role_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleAccessScope_Rules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleAccessScope_Rules) ProtoMessage() {}

func (x *SimpleAccessScope_Rules) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleAccessScope_Rules.ProtoReflect.Descriptor instead.
func (*SimpleAccessScope_Rules) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SimpleAccessScope_Rules) GetIncludedClusters() []string {
	if x != nil {
		return x.IncludedClusters
	}
	return nil
}

func (x *SimpleAccessScope_Rules) GetIncludedNamespaces() []*SimpleAccessScope_Rules_Namespace {
	if x != nil {
		return x.IncludedNamespaces
	}
	return nil
}

func (x *SimpleAccessScope_Rules) GetClusterLabelSelectors() []*SetBasedLabelSelector {
	if x != nil {
		return x.ClusterLabelSelectors
	}
	return nil
}

func (x *SimpleAccessScope_Rules) GetNamespaceLabelSelectors() []*SetBasedLabelSelector {
	if x != nil {
		return x.NamespaceLabelSelectors
	}
	return nil
}

type SimpleAccessScope_Rules_Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Both fields must be set.
	ClusterName   string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	NamespaceName string `protobuf:"bytes,2,opt,name=namespace_name,json=namespaceName,proto3" json:"namespace_name,omitempty"`
}

func (x *SimpleAccessScope_Rules_Namespace) Reset() {
	*x = SimpleAccessScope_Rules_Namespace{}
	mi := &file_storage_role_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleAccessScope_Rules_Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleAccessScope_Rules_Namespace) ProtoMessage() {}

func (x *SimpleAccessScope_Rules_Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleAccessScope_Rules_Namespace.ProtoReflect.Descriptor instead.
func (*SimpleAccessScope_Rules_Namespace) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *SimpleAccessScope_Rules_Namespace) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *SimpleAccessScope_Rules_Namespace) GetNamespaceName() string {
	if x != nil {
		return x.NamespaceName
	}
	return ""
}

type EffectiveAccessScope_Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State  EffectiveAccessScope_State `protobuf:"varint,3,opt,name=state,proto3,enum=storage.EffectiveAccessScope_State" json:"state,omitempty"`
	Labels map[string]string          `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EffectiveAccessScope_Namespace) Reset() {
	*x = EffectiveAccessScope_Namespace{}
	mi := &file_storage_role_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EffectiveAccessScope_Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveAccessScope_Namespace) ProtoMessage() {}

func (x *EffectiveAccessScope_Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveAccessScope_Namespace.ProtoReflect.Descriptor instead.
func (*EffectiveAccessScope_Namespace) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{3, 0}
}

func (x *EffectiveAccessScope_Namespace) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EffectiveAccessScope_Namespace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EffectiveAccessScope_Namespace) GetState() EffectiveAccessScope_State {
	if x != nil {
		return x.State
	}
	return EffectiveAccessScope_UNKNOWN
}

func (x *EffectiveAccessScope_Namespace) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type EffectiveAccessScope_Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State      EffectiveAccessScope_State        `protobuf:"varint,3,opt,name=state,proto3,enum=storage.EffectiveAccessScope_State" json:"state,omitempty"`
	Labels     map[string]string                 `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Namespaces []*EffectiveAccessScope_Namespace `protobuf:"bytes,4,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *EffectiveAccessScope_Cluster) Reset() {
	*x = EffectiveAccessScope_Cluster{}
	mi := &file_storage_role_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EffectiveAccessScope_Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveAccessScope_Cluster) ProtoMessage() {}

func (x *EffectiveAccessScope_Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_storage_role_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveAccessScope_Cluster.ProtoReflect.Descriptor instead.
func (*EffectiveAccessScope_Cluster) Descriptor() ([]byte, []int) {
	return file_storage_role_proto_rawDescGZIP(), []int{3, 1}
}

func (x *EffectiveAccessScope_Cluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EffectiveAccessScope_Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EffectiveAccessScope_Cluster) GetState() EffectiveAccessScope_State {
	if x != nil {
		return x.State
	}
	return EffectiveAccessScope_UNKNOWN
}

func (x *EffectiveAccessScope_Cluster) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *EffectiveAccessScope_Cluster) GetNamespaces() []*EffectiveAccessScope_Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

var File_storage_role_proto protoreflect.FileDescriptor

var file_storage_role_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x03, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0d,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x55, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x02, 0x18, 0x01, 0x52, 0x10, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a,
	0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x52, 0x06,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a, 0x54, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x04,
	0x10, 0x05, 0x22, 0xb0, 0x02, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x12, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x52, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a,
	0x54, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd9, 0x04, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x52, 0x06, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x1a, 0x9c, 0x03, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x5b, 0x0a, 0x13, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x17, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x5a,
	0x0a, 0x19, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x42,
	0x61, 0x73, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x17, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x55, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xc7, 0x05, 0x0a, 0x14, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xf2, 0x01,
	0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0xb7, 0x02, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x03, 0x2a, 0x3f, 0x0a, 0x06, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x57, 0x52,
	0x49, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x42, 0x2e, 0x0a, 0x19,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_role_proto_rawDescOnce sync.Once
	file_storage_role_proto_rawDescData = file_storage_role_proto_rawDesc
)

func file_storage_role_proto_rawDescGZIP() []byte {
	file_storage_role_proto_rawDescOnce.Do(func() {
		file_storage_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_role_proto_rawDescData)
	})
	return file_storage_role_proto_rawDescData
}

var file_storage_role_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_storage_role_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_storage_role_proto_goTypes = []any{
	(Access)(0),                     // 0: storage.Access
	(EffectiveAccessScope_State)(0), // 1: storage.EffectiveAccessScope.State
	(*Role)(nil),                    // 2: storage.Role
	(*PermissionSet)(nil),           // 3: storage.PermissionSet
	(*SimpleAccessScope)(nil),       // 4: storage.SimpleAccessScope
	(*EffectiveAccessScope)(nil),    // 5: storage.EffectiveAccessScope
	nil,                             // 6: storage.Role.ResourceToAccessEntry
	nil,                             // 7: storage.PermissionSet.ResourceToAccessEntry
	(*SimpleAccessScope_Rules)(nil), // 8: storage.SimpleAccessScope.Rules
	(*SimpleAccessScope_Rules_Namespace)(nil), // 9: storage.SimpleAccessScope.Rules.Namespace
	(*EffectiveAccessScope_Namespace)(nil),    // 10: storage.EffectiveAccessScope.Namespace
	(*EffectiveAccessScope_Cluster)(nil),      // 11: storage.EffectiveAccessScope.Cluster
	nil,                                       // 12: storage.EffectiveAccessScope.Namespace.LabelsEntry
	nil,                                       // 13: storage.EffectiveAccessScope.Cluster.LabelsEntry
	(*Traits)(nil),                            // 14: storage.Traits
	(*SetBasedLabelSelector)(nil),             // 15: storage.SetBasedLabelSelector
}
var file_storage_role_proto_depIdxs = []int32{
	0,  // 0: storage.Role.global_access:type_name -> storage.Access
	6,  // 1: storage.Role.resource_to_access:type_name -> storage.Role.ResourceToAccessEntry
	14, // 2: storage.Role.traits:type_name -> storage.Traits
	7,  // 3: storage.PermissionSet.resource_to_access:type_name -> storage.PermissionSet.ResourceToAccessEntry
	14, // 4: storage.PermissionSet.traits:type_name -> storage.Traits
	8,  // 5: storage.SimpleAccessScope.rules:type_name -> storage.SimpleAccessScope.Rules
	14, // 6: storage.SimpleAccessScope.traits:type_name -> storage.Traits
	11, // 7: storage.EffectiveAccessScope.clusters:type_name -> storage.EffectiveAccessScope.Cluster
	0,  // 8: storage.Role.ResourceToAccessEntry.value:type_name -> storage.Access
	0,  // 9: storage.PermissionSet.ResourceToAccessEntry.value:type_name -> storage.Access
	9,  // 10: storage.SimpleAccessScope.Rules.included_namespaces:type_name -> storage.SimpleAccessScope.Rules.Namespace
	15, // 11: storage.SimpleAccessScope.Rules.cluster_label_selectors:type_name -> storage.SetBasedLabelSelector
	15, // 12: storage.SimpleAccessScope.Rules.namespace_label_selectors:type_name -> storage.SetBasedLabelSelector
	1,  // 13: storage.EffectiveAccessScope.Namespace.state:type_name -> storage.EffectiveAccessScope.State
	12, // 14: storage.EffectiveAccessScope.Namespace.labels:type_name -> storage.EffectiveAccessScope.Namespace.LabelsEntry
	1,  // 15: storage.EffectiveAccessScope.Cluster.state:type_name -> storage.EffectiveAccessScope.State
	13, // 16: storage.EffectiveAccessScope.Cluster.labels:type_name -> storage.EffectiveAccessScope.Cluster.LabelsEntry
	10, // 17: storage.EffectiveAccessScope.Cluster.namespaces:type_name -> storage.EffectiveAccessScope.Namespace
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_storage_role_proto_init() }
func file_storage_role_proto_init() {
	if File_storage_role_proto != nil {
		return
	}
	file_storage_labels_proto_init()
	file_storage_traits_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_role_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_role_proto_goTypes,
		DependencyIndexes: file_storage_role_proto_depIdxs,
		EnumInfos:         file_storage_role_proto_enumTypes,
		MessageInfos:      file_storage_role_proto_msgTypes,
	}.Build()
	File_storage_role_proto = out.File
	file_storage_role_proto_rawDesc = nil
	file_storage_role_proto_goTypes = nil
	file_storage_role_proto_depIdxs = nil
}
