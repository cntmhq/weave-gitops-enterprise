//
// This file holds the protobuf definitions for messages and enums
// used in the Weave GitOps GitOpsSets API.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/gitopssets/types.proto

package api

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

type GitOpsSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace          string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Inventory          []*ResourceRef    `protobuf:"bytes,3,rep,name=inventory,proto3" json:"inventory,omitempty"`
	Conditions         []*Condition      `protobuf:"bytes,4,rep,name=conditions,proto3" json:"conditions,omitempty"`
	ClusterName        string            `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Type               string            `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Labels             map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations        map[string]string `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Suspended          bool              `protobuf:"varint,9,opt,name=suspended,proto3" json:"suspended,omitempty"`
	ObservedGeneration int64             `protobuf:"varint,10,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	Yaml               string            `protobuf:"bytes,11,opt,name=yaml,proto3" json:"yaml,omitempty"`
	ObjectRef          *ObjectRef        `protobuf:"bytes,12,opt,name=object_ref,json=objectRef,proto3" json:"object_ref,omitempty"`
}

func (x *GitOpsSet) Reset() {
	*x = GitOpsSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitOpsSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitOpsSet) ProtoMessage() {}

func (x *GitOpsSet) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitOpsSet.ProtoReflect.Descriptor instead.
func (*GitOpsSet) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{0}
}

func (x *GitOpsSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GitOpsSet) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GitOpsSet) GetInventory() []*ResourceRef {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *GitOpsSet) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *GitOpsSet) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GitOpsSet) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GitOpsSet) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *GitOpsSet) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *GitOpsSet) GetSuspended() bool {
	if x != nil {
		return x.Suspended
	}
	return false
}

func (x *GitOpsSet) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *GitOpsSet) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

func (x *GitOpsSet) GetObjectRef() *ObjectRef {
	if x != nil {
		return x.ObjectRef
	}
	return nil
}

type GitOpsSetListError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace   string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *GitOpsSetListError) Reset() {
	*x = GitOpsSetListError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitOpsSetListError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitOpsSetListError) ProtoMessage() {}

func (x *GitOpsSetListError) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitOpsSetListError.ProtoReflect.Descriptor instead.
func (*GitOpsSetListError) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{1}
}

func (x *GitOpsSetListError) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GitOpsSetListError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GitOpsSetListError) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type ResourceRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ResourceRef) Reset() {
	*x = ResourceRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRef) ProtoMessage() {}

func (x *ResourceRef) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRef.ProtoReflect.Descriptor instead.
func (*ResourceRef) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceRef) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceRef) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Reason    string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	Message   string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp string `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{3}
}

func (x *Condition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Condition) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Condition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Condition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Condition) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type GitOpsSetGenerator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List          string `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
	GitRepository string `protobuf:"bytes,2,opt,name=git_repository,json=gitRepository,proto3" json:"git_repository,omitempty"`
}

func (x *GitOpsSetGenerator) Reset() {
	*x = GitOpsSetGenerator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitOpsSetGenerator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitOpsSetGenerator) ProtoMessage() {}

func (x *GitOpsSetGenerator) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitOpsSetGenerator.ProtoReflect.Descriptor instead.
func (*GitOpsSetGenerator) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{4}
}

func (x *GitOpsSetGenerator) GetList() string {
	if x != nil {
		return x.List
	}
	return ""
}

func (x *GitOpsSetGenerator) GetGitRepository() string {
	if x != nil {
		return x.GitRepository
	}
	return ""
}

// GroupVersionKind represents an objects Kubernetes API type data
type GroupVersionKind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Kind    string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GroupVersionKind) Reset() {
	*x = GroupVersionKind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupVersionKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupVersionKind) ProtoMessage() {}

func (x *GroupVersionKind) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupVersionKind.ProtoReflect.Descriptor instead.
func (*GroupVersionKind) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{5}
}

func (x *GroupVersionKind) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GroupVersionKind) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *GroupVersionKind) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload     string              `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	ClusterName string              `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Tenant      string              `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Uid         string              `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Inventory   []*GroupVersionKind `protobuf:"bytes,5,rep,name=inventory,proto3" json:"inventory,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{6}
}

func (x *Object) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Object) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Object) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *Object) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Object) GetInventory() []*GroupVersionKind {
	if x != nil {
		return x.Inventory
	}
	return nil
}

type ObjectRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind        string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ClusterName string `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *ObjectRef) Reset() {
	*x = ObjectRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gitopssets_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRef) ProtoMessage() {}

func (x *ObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_api_gitopssets_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRef.ProtoReflect.Descriptor instead.
func (*ObjectRef) Descriptor() ([]byte, []int) {
	return file_api_gitopssets_types_proto_rawDescGZIP(), []int{7}
}

func (x *ObjectRef) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ObjectRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ObjectRef) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

var File_api_gitopssets_types_proto protoreflect.FileDescriptor

var file_api_gitopssets_types_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x69,
	0x74, 0x6f, 0x70, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x8a, 0x05, 0x0a, 0x09,
	0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x53, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x53, 0x65, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x4b, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73,
	0x53, 0x65, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12,
	0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x79, 0x61, 0x6d, 0x6c, 0x12, 0x37, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72,
	0x65, 0x66, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6f, 0x0a, 0x12, 0x47, 0x69, 0x74, 0x4f,
	0x70, 0x73, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4f, 0x0a, 0x12,
	0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x53, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x69, 0x74, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x67, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x56, 0x0a,
	0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xae, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x69, 0x74,
	0x6f, 0x70, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x09, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x74, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x3e, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2d, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x67, 0x69,
	0x74, 0x6f, 0x70, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_gitopssets_types_proto_rawDescOnce sync.Once
	file_api_gitopssets_types_proto_rawDescData = file_api_gitopssets_types_proto_rawDesc
)

func file_api_gitopssets_types_proto_rawDescGZIP() []byte {
	file_api_gitopssets_types_proto_rawDescOnce.Do(func() {
		file_api_gitopssets_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gitopssets_types_proto_rawDescData)
	})
	return file_api_gitopssets_types_proto_rawDescData
}

var file_api_gitopssets_types_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_gitopssets_types_proto_goTypes = []interface{}{
	(*GitOpsSet)(nil),          // 0: gitopssets.v1.GitOpsSet
	(*GitOpsSetListError)(nil), // 1: gitopssets.v1.GitOpsSetListError
	(*ResourceRef)(nil),        // 2: gitopssets.v1.ResourceRef
	(*Condition)(nil),          // 3: gitopssets.v1.Condition
	(*GitOpsSetGenerator)(nil), // 4: gitopssets.v1.GitOpsSetGenerator
	(*GroupVersionKind)(nil),   // 5: gitopssets.v1.GroupVersionKind
	(*Object)(nil),             // 6: gitopssets.v1.Object
	(*ObjectRef)(nil),          // 7: gitopssets.v1.ObjectRef
	nil,                        // 8: gitopssets.v1.GitOpsSet.LabelsEntry
	nil,                        // 9: gitopssets.v1.GitOpsSet.AnnotationsEntry
}
var file_api_gitopssets_types_proto_depIdxs = []int32{
	2, // 0: gitopssets.v1.GitOpsSet.inventory:type_name -> gitopssets.v1.ResourceRef
	3, // 1: gitopssets.v1.GitOpsSet.conditions:type_name -> gitopssets.v1.Condition
	8, // 2: gitopssets.v1.GitOpsSet.labels:type_name -> gitopssets.v1.GitOpsSet.LabelsEntry
	9, // 3: gitopssets.v1.GitOpsSet.annotations:type_name -> gitopssets.v1.GitOpsSet.AnnotationsEntry
	7, // 4: gitopssets.v1.GitOpsSet.object_ref:type_name -> gitopssets.v1.ObjectRef
	5, // 5: gitopssets.v1.Object.inventory:type_name -> gitopssets.v1.GroupVersionKind
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_gitopssets_types_proto_init() }
func file_api_gitopssets_types_proto_init() {
	if File_api_gitopssets_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gitopssets_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitOpsSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_gitopssets_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitOpsSetListError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_gitopssets_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_gitopssets_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_gitopssets_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitOpsSetGenerator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_gitopssets_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupVersionKind); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_gitopssets_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_gitopssets_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_gitopssets_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_gitopssets_types_proto_goTypes,
		DependencyIndexes: file_api_gitopssets_types_proto_depIdxs,
		MessageInfos:      file_api_gitopssets_types_proto_msgTypes,
	}.Build()
	File_api_gitopssets_types_proto = out.File
	file_api_gitopssets_types_proto_rawDesc = nil
	file_api_gitopssets_types_proto_goTypes = nil
	file_api_gitopssets_types_proto_depIdxs = nil
}
