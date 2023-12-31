// Copyright 2021 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: bp2build_metrics.proto

package bp2build_metrics_proto

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

type UnconvertedReasonType int32

const (
	// Bp2build does not know how to convert this specific module for some reason
	// not covered by other reason types. The reason detail should explain the
	// specific issue.
	UnconvertedReasonType_UNSUPPORTED UnconvertedReasonType = 0
	// The module was already defined in a BUILD file available in the source tree.
	UnconvertedReasonType_DEFINED_IN_BUILD_FILE UnconvertedReasonType = 1
	// The module was explicitly denylisted by name.
	UnconvertedReasonType_DENYLISTED UnconvertedReasonType = 2
	// The module's type has no bp2build implementation.
	UnconvertedReasonType_TYPE_UNSUPPORTED UnconvertedReasonType = 3
	// The module has a property not yet supported. The detail field should
	// name the unsupported property name.
	UnconvertedReasonType_PROPERTY_UNSUPPORTED UnconvertedReasonType = 4
	// The module has an unconverted dependency. The detail should consist of
	// the name of the unconverted module.
	UnconvertedReasonType_UNCONVERTED_DEP UnconvertedReasonType = 5
	// The module has a source file with the same name as the module itself.
	UnconvertedReasonType_SRC_NAME_COLLISION UnconvertedReasonType = 6
)

// Enum value maps for UnconvertedReasonType.
var (
	UnconvertedReasonType_name = map[int32]string{
		0: "UNSUPPORTED",
		1: "DEFINED_IN_BUILD_FILE",
		2: "DENYLISTED",
		3: "TYPE_UNSUPPORTED",
		4: "PROPERTY_UNSUPPORTED",
		5: "UNCONVERTED_DEP",
		6: "SRC_NAME_COLLISION",
	}
	UnconvertedReasonType_value = map[string]int32{
		"UNSUPPORTED":           0,
		"DEFINED_IN_BUILD_FILE": 1,
		"DENYLISTED":            2,
		"TYPE_UNSUPPORTED":      3,
		"PROPERTY_UNSUPPORTED":  4,
		"UNCONVERTED_DEP":       5,
		"SRC_NAME_COLLISION":    6,
	}
)

func (x UnconvertedReasonType) Enum() *UnconvertedReasonType {
	p := new(UnconvertedReasonType)
	*p = x
	return p
}

func (x UnconvertedReasonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnconvertedReasonType) Descriptor() protoreflect.EnumDescriptor {
	return file_bp2build_metrics_proto_enumTypes[0].Descriptor()
}

func (UnconvertedReasonType) Type() protoreflect.EnumType {
	return &file_bp2build_metrics_proto_enumTypes[0]
}

func (x UnconvertedReasonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UnconvertedReasonType.Descriptor instead.
func (UnconvertedReasonType) EnumDescriptor() ([]byte, []int) {
	return file_bp2build_metrics_proto_rawDescGZIP(), []int{0}
}

type Bp2BuildMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of Soong modules converted to generated targets
	GeneratedModuleCount uint64 `protobuf:"varint,1,opt,name=generatedModuleCount,proto3" json:"generatedModuleCount,omitempty"`
	// Total number of Soong modules converted to handcrafted targets
	HandCraftedModuleCount uint64 `protobuf:"varint,2,opt,name=handCraftedModuleCount,proto3" json:"handCraftedModuleCount,omitempty"`
	// Total number of unconverted Soong modules
	UnconvertedModuleCount uint64 `protobuf:"varint,3,opt,name=unconvertedModuleCount,proto3" json:"unconvertedModuleCount,omitempty"`
	// Counts of symlinks in synthetic bazel workspace
	WorkspaceSymlinkCount uint64 `protobuf:"varint,9,opt,name=workspaceSymlinkCount,proto3" json:"workspaceSymlinkCount,omitempty"`
	// Counts of mkdir calls during creation of synthetic bazel workspace
	WorkspaceMkDirCount uint64 `protobuf:"varint,10,opt,name=workspaceMkDirCount,proto3" json:"workspaceMkDirCount,omitempty"`
	// Counts of generated Bazel targets per Bazel rule class
	RuleClassCount map[string]uint64 `protobuf:"bytes,4,rep,name=ruleClassCount,proto3" json:"ruleClassCount,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// List of converted modules
	ConvertedModules []string `protobuf:"bytes,5,rep,name=convertedModules,proto3" json:"convertedModules,omitempty"`
	// Unconverted modules, mapped to the reason the module was not converted.
	UnconvertedModules map[string]*UnconvertedReason `protobuf:"bytes,11,rep,name=unconvertedModules,proto3" json:"unconvertedModules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Counts of converted modules by module type.
	ConvertedModuleTypeCount map[string]uint64 `protobuf:"bytes,6,rep,name=convertedModuleTypeCount,proto3" json:"convertedModuleTypeCount,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Counts of total modules by module type.
	TotalModuleTypeCount map[string]uint64 `protobuf:"bytes,7,rep,name=totalModuleTypeCount,proto3" json:"totalModuleTypeCount,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// List of traced runtime events of bp2build, useful for tracking bp2build
	// runtime.
	Events []*Event `protobuf:"bytes,8,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Bp2BuildMetrics) Reset() {
	*x = Bp2BuildMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bp2build_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bp2BuildMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bp2BuildMetrics) ProtoMessage() {}

func (x *Bp2BuildMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_bp2build_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bp2BuildMetrics.ProtoReflect.Descriptor instead.
func (*Bp2BuildMetrics) Descriptor() ([]byte, []int) {
	return file_bp2build_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Bp2BuildMetrics) GetGeneratedModuleCount() uint64 {
	if x != nil {
		return x.GeneratedModuleCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetHandCraftedModuleCount() uint64 {
	if x != nil {
		return x.HandCraftedModuleCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetUnconvertedModuleCount() uint64 {
	if x != nil {
		return x.UnconvertedModuleCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetWorkspaceSymlinkCount() uint64 {
	if x != nil {
		return x.WorkspaceSymlinkCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetWorkspaceMkDirCount() uint64 {
	if x != nil {
		return x.WorkspaceMkDirCount
	}
	return 0
}

func (x *Bp2BuildMetrics) GetRuleClassCount() map[string]uint64 {
	if x != nil {
		return x.RuleClassCount
	}
	return nil
}

func (x *Bp2BuildMetrics) GetConvertedModules() []string {
	if x != nil {
		return x.ConvertedModules
	}
	return nil
}

func (x *Bp2BuildMetrics) GetUnconvertedModules() map[string]*UnconvertedReason {
	if x != nil {
		return x.UnconvertedModules
	}
	return nil
}

func (x *Bp2BuildMetrics) GetConvertedModuleTypeCount() map[string]uint64 {
	if x != nil {
		return x.ConvertedModuleTypeCount
	}
	return nil
}

func (x *Bp2BuildMetrics) GetTotalModuleTypeCount() map[string]uint64 {
	if x != nil {
		return x.TotalModuleTypeCount
	}
	return nil
}

func (x *Bp2BuildMetrics) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// Traced runtime event of bp2build.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The event name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The absolute start time of the event
	// The number of nanoseconds elapsed since January 1, 1970 UTC.
	StartTime uint64 `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The real running time.
	// The number of nanoseconds elapsed since start_time.
	RealTime uint64 `protobuf:"varint,3,opt,name=real_time,json=realTime,proto3" json:"real_time,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bp2build_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_bp2build_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_bp2build_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetStartTime() uint64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Event) GetRealTime() uint64 {
	if x != nil {
		return x.RealTime
	}
	return 0
}

type UnconvertedReason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of reason that the module could not be converted.
	Type UnconvertedReasonType `protobuf:"varint,1,opt,name=type,proto3,enum=soong_build_bp2build_metrics.UnconvertedReasonType" json:"type,omitempty"`
	// Descriptive details describing why the module could not be converted.
	// This detail should be kept very short and should be in the context of
	// the type. (Otherwise, this would significantly bloat metrics.)
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *UnconvertedReason) Reset() {
	*x = UnconvertedReason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bp2build_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnconvertedReason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnconvertedReason) ProtoMessage() {}

func (x *UnconvertedReason) ProtoReflect() protoreflect.Message {
	mi := &file_bp2build_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnconvertedReason.ProtoReflect.Descriptor instead.
func (*UnconvertedReason) Descriptor() ([]byte, []int) {
	return file_bp2build_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *UnconvertedReason) GetType() UnconvertedReasonType {
	if x != nil {
		return x.Type
	}
	return UnconvertedReasonType_UNSUPPORTED
}

func (x *UnconvertedReason) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

var File_bp2build_metrics_proto protoreflect.FileDescriptor

var file_bp2build_metrics_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xc0, 0x09, 0x0a, 0x0f, 0x42, 0x70, 0x32, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36,
	0x0a, 0x16, 0x68, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16,
	0x68, 0x61, 0x6e, 0x64, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34,
	0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x79, 0x6d, 0x6c, 0x69,
	0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x79, 0x6d, 0x6c, 0x69, 0x6e, 0x6b, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x4d, 0x6b, 0x44, 0x69, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x6b, 0x44, 0x69,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x69, 0x0a, 0x0e, 0x72, 0x75, 0x6c, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x42, 0x70,
	0x32, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x52, 0x75,
	0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0e, 0x72, 0x75, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x75, 0x0a,
	0x12, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x73, 0x6f, 0x6f, 0x6e,
	0x67, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x42, 0x70, 0x32, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x12, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x87, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x42, 0x70, 0x32, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x7b,
	0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x73,
	0x6f, 0x6f, 0x6e, 0x67, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x42, 0x70, 0x32, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f,
	0x6f, 0x6e, 0x67, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x52, 0x75, 0x6c, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x76, 0x0a, 0x17, 0x55,
	0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x47, 0x0a, 0x19, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x74, 0x0a, 0x11, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2a, 0xb0, 0x01, 0x0a, 0x15, 0x55, 0x6e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x49,
	0x4e, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x44, 0x45, 0x4e, 0x59, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45,
	0x50, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x52, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x43, 0x4f, 0x4c, 0x4c, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x42, 0x31, 0x5a, 0x2f, 0x61,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2f, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x2f, 0x75, 0x69, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x70, 0x32, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bp2build_metrics_proto_rawDescOnce sync.Once
	file_bp2build_metrics_proto_rawDescData = file_bp2build_metrics_proto_rawDesc
)

func file_bp2build_metrics_proto_rawDescGZIP() []byte {
	file_bp2build_metrics_proto_rawDescOnce.Do(func() {
		file_bp2build_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_bp2build_metrics_proto_rawDescData)
	})
	return file_bp2build_metrics_proto_rawDescData
}

var file_bp2build_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bp2build_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bp2build_metrics_proto_goTypes = []interface{}{
	(UnconvertedReasonType)(0), // 0: soong_build_bp2build_metrics.UnconvertedReasonType
	(*Bp2BuildMetrics)(nil),    // 1: soong_build_bp2build_metrics.Bp2BuildMetrics
	(*Event)(nil),              // 2: soong_build_bp2build_metrics.Event
	(*UnconvertedReason)(nil),  // 3: soong_build_bp2build_metrics.UnconvertedReason
	nil,                        // 4: soong_build_bp2build_metrics.Bp2BuildMetrics.RuleClassCountEntry
	nil,                        // 5: soong_build_bp2build_metrics.Bp2BuildMetrics.UnconvertedModulesEntry
	nil,                        // 6: soong_build_bp2build_metrics.Bp2BuildMetrics.ConvertedModuleTypeCountEntry
	nil,                        // 7: soong_build_bp2build_metrics.Bp2BuildMetrics.TotalModuleTypeCountEntry
}
var file_bp2build_metrics_proto_depIdxs = []int32{
	4, // 0: soong_build_bp2build_metrics.Bp2BuildMetrics.ruleClassCount:type_name -> soong_build_bp2build_metrics.Bp2BuildMetrics.RuleClassCountEntry
	5, // 1: soong_build_bp2build_metrics.Bp2BuildMetrics.unconvertedModules:type_name -> soong_build_bp2build_metrics.Bp2BuildMetrics.UnconvertedModulesEntry
	6, // 2: soong_build_bp2build_metrics.Bp2BuildMetrics.convertedModuleTypeCount:type_name -> soong_build_bp2build_metrics.Bp2BuildMetrics.ConvertedModuleTypeCountEntry
	7, // 3: soong_build_bp2build_metrics.Bp2BuildMetrics.totalModuleTypeCount:type_name -> soong_build_bp2build_metrics.Bp2BuildMetrics.TotalModuleTypeCountEntry
	2, // 4: soong_build_bp2build_metrics.Bp2BuildMetrics.events:type_name -> soong_build_bp2build_metrics.Event
	0, // 5: soong_build_bp2build_metrics.UnconvertedReason.type:type_name -> soong_build_bp2build_metrics.UnconvertedReasonType
	3, // 6: soong_build_bp2build_metrics.Bp2BuildMetrics.UnconvertedModulesEntry.value:type_name -> soong_build_bp2build_metrics.UnconvertedReason
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_bp2build_metrics_proto_init() }
func file_bp2build_metrics_proto_init() {
	if File_bp2build_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bp2build_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bp2BuildMetrics); i {
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
		file_bp2build_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_bp2build_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnconvertedReason); i {
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
			RawDescriptor: file_bp2build_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bp2build_metrics_proto_goTypes,
		DependencyIndexes: file_bp2build_metrics_proto_depIdxs,
		EnumInfos:         file_bp2build_metrics_proto_enumTypes,
		MessageInfos:      file_bp2build_metrics_proto_msgTypes,
	}.Build()
	File_bp2build_metrics_proto = out.File
	file_bp2build_metrics_proto_rawDesc = nil
	file_bp2build_metrics_proto_goTypes = nil
	file_bp2build_metrics_proto_depIdxs = nil
}
