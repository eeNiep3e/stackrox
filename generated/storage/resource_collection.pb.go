// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: storage/resource_collection.proto

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

type MatchType int32

const (
	MatchType_EXACT MatchType = 0
	MatchType_REGEX MatchType = 1
)

// Enum value maps for MatchType.
var (
	MatchType_name = map[int32]string{
		0: "EXACT",
		1: "REGEX",
	}
	MatchType_value = map[string]int32{
		"EXACT": 0,
		"REGEX": 1,
	}
)

func (x MatchType) Enum() *MatchType {
	p := new(MatchType)
	*p = x
	return p
}

func (x MatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_resource_collection_proto_enumTypes[0].Descriptor()
}

func (MatchType) Type() protoreflect.EnumType {
	return &file_storage_resource_collection_proto_enumTypes[0]
}

func (x MatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchType.Descriptor instead.
func (MatchType) EnumDescriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{0}
}

type ResourceCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Collection ID" sql:"pk"`     // @gotags: search:"Collection ID" sql:"pk"
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Collection Name" sql:"unique"` // @gotags: search:"Collection Name" sql:"unique"
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	CreatedBy   *SlimUser              `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty" sql:"ignore_labels(User ID)"` // @gotags: sql:"ignore_labels(User ID)"
	UpdatedBy   *SlimUser              `protobuf:"bytes,7,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty" sql:"ignore_labels(User ID)"` // @gotags: sql:"ignore_labels(User ID)"
	// `resource_selectors` resolve as disjunction (OR) with each-other and with selectors from `embedded_collections`. For MVP, the size of resource_selectors will at most be 1 from UX standpoint.
	ResourceSelectors   []*ResourceSelector                              `protobuf:"bytes,8,rep,name=resource_selectors,json=resourceSelectors,proto3" json:"resource_selectors,omitempty"`
	EmbeddedCollections []*ResourceCollection_EmbeddedResourceCollection `protobuf:"bytes,9,rep,name=embedded_collections,json=embeddedCollections,proto3" json:"embedded_collections,omitempty"`
}

func (x *ResourceCollection) Reset() {
	*x = ResourceCollection{}
	mi := &file_storage_resource_collection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceCollection) ProtoMessage() {}

func (x *ResourceCollection) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceCollection.ProtoReflect.Descriptor instead.
func (*ResourceCollection) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceCollection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceCollection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceCollection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ResourceCollection) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ResourceCollection) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *ResourceCollection) GetCreatedBy() *SlimUser {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ResourceCollection) GetUpdatedBy() *SlimUser {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *ResourceCollection) GetResourceSelectors() []*ResourceSelector {
	if x != nil {
		return x.ResourceSelectors
	}
	return nil
}

func (x *ResourceCollection) GetEmbeddedCollections() []*ResourceCollection_EmbeddedResourceCollection {
	if x != nil {
		return x.EmbeddedCollections
	}
	return nil
}

type ResourceSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `rules` resolve as a conjunction (AND).
	Rules []*SelectorRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ResourceSelector) Reset() {
	*x = ResourceSelector{}
	mi := &file_storage_resource_collection_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSelector) ProtoMessage() {}

func (x *ResourceSelector) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSelector.ProtoReflect.Descriptor instead.
func (*ResourceSelector) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceSelector) GetRules() []*SelectorRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type SelectorRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `field_name` can be one of the following:
	// - Cluster
	// - Cluster Label
	// - Namespace
	// - Namespace Label
	// - Namespace Annotation
	// - Deployment
	// - Deployment Label
	// - Deployment Annotation
	FieldName string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// 'operator' only supports disjunction (OR) currently
	Operator BooleanOperator `protobuf:"varint,2,opt,name=operator,proto3,enum=storage.BooleanOperator" json:"operator,omitempty"`
	// `values` resolve as a conjunction (AND) or disjunction (OR) depending on operator. For MVP, only OR is supported from UX standpoint.
	Values []*RuleValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *SelectorRule) Reset() {
	*x = SelectorRule{}
	mi := &file_storage_resource_collection_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SelectorRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectorRule) ProtoMessage() {}

func (x *SelectorRule) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectorRule.ProtoReflect.Descriptor instead.
func (*SelectorRule) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{2}
}

func (x *SelectorRule) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *SelectorRule) GetOperator() BooleanOperator {
	if x != nil {
		return x.Operator
	}
	return BooleanOperator_OR
}

func (x *SelectorRule) GetValues() []*RuleValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type RuleValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     string    `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	MatchType MatchType `protobuf:"varint,2,opt,name=match_type,json=matchType,proto3,enum=storage.MatchType" json:"match_type,omitempty"`
}

func (x *RuleValue) Reset() {
	*x = RuleValue{}
	mi := &file_storage_resource_collection_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleValue) ProtoMessage() {}

func (x *RuleValue) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleValue.ProtoReflect.Descriptor instead.
func (*RuleValue) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{3}
}

func (x *RuleValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RuleValue) GetMatchType() MatchType {
	if x != nil {
		return x.MatchType
	}
	return MatchType_EXACT
}

type ResourceCollection_EmbeddedResourceCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 'id' is searchable to force a separate table
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Embedded Collection ID" sql:"fk(ResourceCollection:id),restrict-delete"` // @gotags: search:"Embedded Collection ID" sql:"fk(ResourceCollection:id),restrict-delete"
}

func (x *ResourceCollection_EmbeddedResourceCollection) Reset() {
	*x = ResourceCollection_EmbeddedResourceCollection{}
	mi := &file_storage_resource_collection_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceCollection_EmbeddedResourceCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceCollection_EmbeddedResourceCollection) ProtoMessage() {}

func (x *ResourceCollection_EmbeddedResourceCollection) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceCollection_EmbeddedResourceCollection.ProtoReflect.Descriptor instead.
func (*ResourceCollection_EmbeddedResourceCollection) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResourceCollection_EmbeddedResourceCollection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_storage_resource_collection_proto protoreflect.FileDescriptor

var file_storage_resource_collection_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x04, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x30,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6c, 0x69,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x30, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53,
	0x6c, 0x69, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x48, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x69, 0x0a, 0x14,
	0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x13, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x2c, 0x0a, 0x1a, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x09, 0x52, 0x75, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x21,
	0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x58, 0x41, 0x43, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x47, 0x45, 0x58, 0x10,
	0x01, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11,
	0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_resource_collection_proto_rawDescOnce sync.Once
	file_storage_resource_collection_proto_rawDescData = file_storage_resource_collection_proto_rawDesc
)

func file_storage_resource_collection_proto_rawDescGZIP() []byte {
	file_storage_resource_collection_proto_rawDescOnce.Do(func() {
		file_storage_resource_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_resource_collection_proto_rawDescData)
	})
	return file_storage_resource_collection_proto_rawDescData
}

var file_storage_resource_collection_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_resource_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_storage_resource_collection_proto_goTypes = []any{
	(MatchType)(0),             // 0: storage.MatchType
	(*ResourceCollection)(nil), // 1: storage.ResourceCollection
	(*ResourceSelector)(nil),   // 2: storage.ResourceSelector
	(*SelectorRule)(nil),       // 3: storage.SelectorRule
	(*RuleValue)(nil),          // 4: storage.RuleValue
	(*ResourceCollection_EmbeddedResourceCollection)(nil), // 5: storage.ResourceCollection.EmbeddedResourceCollection
	(*timestamppb.Timestamp)(nil),                         // 6: google.protobuf.Timestamp
	(*SlimUser)(nil),                                      // 7: storage.SlimUser
	(BooleanOperator)(0),                                  // 8: storage.BooleanOperator
}
var file_storage_resource_collection_proto_depIdxs = []int32{
	6,  // 0: storage.ResourceCollection.created_at:type_name -> google.protobuf.Timestamp
	6,  // 1: storage.ResourceCollection.last_updated:type_name -> google.protobuf.Timestamp
	7,  // 2: storage.ResourceCollection.created_by:type_name -> storage.SlimUser
	7,  // 3: storage.ResourceCollection.updated_by:type_name -> storage.SlimUser
	2,  // 4: storage.ResourceCollection.resource_selectors:type_name -> storage.ResourceSelector
	5,  // 5: storage.ResourceCollection.embedded_collections:type_name -> storage.ResourceCollection.EmbeddedResourceCollection
	3,  // 6: storage.ResourceSelector.rules:type_name -> storage.SelectorRule
	8,  // 7: storage.SelectorRule.operator:type_name -> storage.BooleanOperator
	4,  // 8: storage.SelectorRule.values:type_name -> storage.RuleValue
	0,  // 9: storage.RuleValue.match_type:type_name -> storage.MatchType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_storage_resource_collection_proto_init() }
func file_storage_resource_collection_proto_init() {
	if File_storage_resource_collection_proto != nil {
		return
	}
	file_storage_policy_proto_init()
	file_storage_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_resource_collection_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_resource_collection_proto_goTypes,
		DependencyIndexes: file_storage_resource_collection_proto_depIdxs,
		EnumInfos:         file_storage_resource_collection_proto_enumTypes,
		MessageInfos:      file_storage_resource_collection_proto_msgTypes,
	}.Build()
	File_storage_resource_collection_proto = out.File
	file_storage_resource_collection_proto_rawDesc = nil
	file_storage_resource_collection_proto_goTypes = nil
	file_storage_resource_collection_proto_depIdxs = nil
}
