// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: app/Adapter/In/ApiGrcp/proto/step.proto

package step

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TaskUuid string `protobuf:"bytes,2,opt,name=task_uuid,json=taskUuid,proto3" json:"task_uuid,omitempty"`
	Sentence string `protobuf:"bytes,3,opt,name=sentence,proto3" json:"sentence,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{0}
}

func (x *Step) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Step) GetTaskUuid() string {
	if x != nil {
		return x.TaskUuid
	}
	return ""
}

func (x *Step) GetSentence() string {
	if x != nil {
		return x.Sentence
	}
	return ""
}

type CreateStepParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskUuid string `protobuf:"bytes,1,opt,name=task_uuid,json=taskUuid,proto3" json:"task_uuid,omitempty"`
	Sentence string `protobuf:"bytes,2,opt,name=sentence,proto3" json:"sentence,omitempty"`
}

func (x *CreateStepParams) Reset() {
	*x = CreateStepParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStepParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStepParams) ProtoMessage() {}

func (x *CreateStepParams) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStepParams.ProtoReflect.Descriptor instead.
func (*CreateStepParams) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStepParams) GetTaskUuid() string {
	if x != nil {
		return x.TaskUuid
	}
	return ""
}

func (x *CreateStepParams) GetSentence() string {
	if x != nil {
		return x.Sentence
	}
	return ""
}

type EditableStepParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskUuid *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=task_uuid,json=taskUuid,proto3" json:"task_uuid,omitempty"`
	Sentence *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=sentence,proto3" json:"sentence,omitempty"`
}

func (x *EditableStepParams) Reset() {
	*x = EditableStepParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditableStepParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditableStepParams) ProtoMessage() {}

func (x *EditableStepParams) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditableStepParams.ProtoReflect.Descriptor instead.
func (*EditableStepParams) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{2}
}

func (x *EditableStepParams) GetTaskUuid() *wrapperspb.StringValue {
	if x != nil {
		return x.TaskUuid
	}
	return nil
}

func (x *EditableStepParams) GetSentence() *wrapperspb.StringValue {
	if x != nil {
		return x.Sentence
	}
	return nil
}

type CreateStepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StepParams *CreateStepParams `protobuf:"bytes,1,opt,name=stepParams,proto3" json:"stepParams,omitempty"`
}

func (x *CreateStepRequest) Reset() {
	*x = CreateStepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStepRequest) ProtoMessage() {}

func (x *CreateStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStepRequest.ProtoReflect.Descriptor instead.
func (*CreateStepRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStepRequest) GetStepParams() *CreateStepParams {
	if x != nil {
		return x.StepParams
	}
	return nil
}

type CreateStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step *Step `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *CreateStepResponse) Reset() {
	*x = CreateStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStepResponse) ProtoMessage() {}

func (x *CreateStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStepResponse.ProtoReflect.Descriptor instead.
func (*CreateStepResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{4}
}

func (x *CreateStepResponse) GetStep() *Step {
	if x != nil {
		return x.Step
	}
	return nil
}

type ReadStepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StepUuid string `protobuf:"bytes,1,opt,name=step_uuid,json=stepUuid,proto3" json:"step_uuid,omitempty"`
}

func (x *ReadStepRequest) Reset() {
	*x = ReadStepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadStepRequest) ProtoMessage() {}

func (x *ReadStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadStepRequest.ProtoReflect.Descriptor instead.
func (*ReadStepRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{5}
}

func (x *ReadStepRequest) GetStepUuid() string {
	if x != nil {
		return x.StepUuid
	}
	return ""
}

type ReadStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step *Step `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *ReadStepResponse) Reset() {
	*x = ReadStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadStepResponse) ProtoMessage() {}

func (x *ReadStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadStepResponse.ProtoReflect.Descriptor instead.
func (*ReadStepResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{6}
}

func (x *ReadStepResponse) GetStep() *Step {
	if x != nil {
		return x.Step
	}
	return nil
}

type UpdateStepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StepUuid   string              `protobuf:"bytes,1,opt,name=step_uuid,json=stepUuid,proto3" json:"step_uuid,omitempty"`
	StepParams *EditableStepParams `protobuf:"bytes,2,opt,name=stepParams,proto3" json:"stepParams,omitempty"`
}

func (x *UpdateStepRequest) Reset() {
	*x = UpdateStepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStepRequest) ProtoMessage() {}

func (x *UpdateStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStepRequest.ProtoReflect.Descriptor instead.
func (*UpdateStepRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStepRequest) GetStepUuid() string {
	if x != nil {
		return x.StepUuid
	}
	return ""
}

func (x *UpdateStepRequest) GetStepParams() *EditableStepParams {
	if x != nil {
		return x.StepParams
	}
	return nil
}

type UpdateStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStepResponse) Reset() {
	*x = UpdateStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStepResponse) ProtoMessage() {}

func (x *UpdateStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStepResponse.ProtoReflect.Descriptor instead.
func (*UpdateStepResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{8}
}

type DeleteStepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StepUuid string `protobuf:"bytes,1,opt,name=step_uuid,json=stepUuid,proto3" json:"step_uuid,omitempty"`
}

func (x *DeleteStepRequest) Reset() {
	*x = DeleteStepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStepRequest) ProtoMessage() {}

func (x *DeleteStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStepRequest.ProtoReflect.Descriptor instead.
func (*DeleteStepRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteStepRequest) GetStepUuid() string {
	if x != nil {
		return x.StepUuid
	}
	return ""
}

type DeleteStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStepResponse) Reset() {
	*x = DeleteStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStepResponse) ProtoMessage() {}

func (x *DeleteStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStepResponse.ProtoReflect.Descriptor instead.
func (*DeleteStepResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{10}
}

type ListStepsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStepsRequest) Reset() {
	*x = ListStepsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStepsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStepsRequest) ProtoMessage() {}

func (x *ListStepsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStepsRequest.ProtoReflect.Descriptor instead.
func (*ListStepsRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{11}
}

type ListStepsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steps []*Step `protobuf:"bytes,1,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *ListStepsResponse) Reset() {
	*x = ListStepsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStepsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStepsResponse) ProtoMessage() {}

func (x *ListStepsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStepsResponse.ProtoReflect.Descriptor instead.
func (*ListStepsResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP(), []int{12}
}

func (x *ListStepsResponse) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

var File_app_Adapter_In_ApiGrcp_proto_step_proto protoreflect.FileDescriptor

var file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x70, 0x2f, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x49, 0x6e,
	0x2f, 0x41, 0x70, 0x69, 0x47, 0x72, 0x63, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x74, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x74, 0x65, 0x70, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x53, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x61, 0x73, 0x6b, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x22, 0x4b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0x89, 0x01, 0x0a, 0x12, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x4b, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a,
	0x73, 0x74, 0x65, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70,
	0x22, 0x2e, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x55, 0x75, 0x69, 0x64,
	0x22, 0x32, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x22, 0x6a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65,
	0x70, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x65, 0x70, 0x55, 0x75, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x65,
	0x70, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x74, 0x65, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x74, 0x65, 0x70, 0x55, 0x75, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x65, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x35, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x65, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x32, 0xc9, 0x02, 0x0a, 0x0b, 0x53, 0x74,
	0x65, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x74, 0x65, 0x70, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x74, 0x65, 0x70, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x65, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73,
	0x74, 0x65, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x65, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x74, 0x65, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x65, 0x70, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x65, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x74, 0x65, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x65, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescOnce sync.Once
	file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescData = file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDesc
)

func file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescGZIP() []byte {
	file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescOnce.Do(func() {
		file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescData)
	})
	return file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDescData
}

var file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_app_Adapter_In_ApiGrcp_proto_step_proto_goTypes = []interface{}{
	(*Step)(nil),                   // 0: step.Step
	(*CreateStepParams)(nil),       // 1: step.CreateStepParams
	(*EditableStepParams)(nil),     // 2: step.EditableStepParams
	(*CreateStepRequest)(nil),      // 3: step.CreateStepRequest
	(*CreateStepResponse)(nil),     // 4: step.CreateStepResponse
	(*ReadStepRequest)(nil),        // 5: step.ReadStepRequest
	(*ReadStepResponse)(nil),       // 6: step.ReadStepResponse
	(*UpdateStepRequest)(nil),      // 7: step.UpdateStepRequest
	(*UpdateStepResponse)(nil),     // 8: step.UpdateStepResponse
	(*DeleteStepRequest)(nil),      // 9: step.DeleteStepRequest
	(*DeleteStepResponse)(nil),     // 10: step.DeleteStepResponse
	(*ListStepsRequest)(nil),       // 11: step.ListStepsRequest
	(*ListStepsResponse)(nil),      // 12: step.ListStepsResponse
	(*wrapperspb.StringValue)(nil), // 13: google.protobuf.StringValue
}
var file_app_Adapter_In_ApiGrcp_proto_step_proto_depIdxs = []int32{
	13, // 0: step.EditableStepParams.task_uuid:type_name -> google.protobuf.StringValue
	13, // 1: step.EditableStepParams.sentence:type_name -> google.protobuf.StringValue
	1,  // 2: step.CreateStepRequest.stepParams:type_name -> step.CreateStepParams
	0,  // 3: step.CreateStepResponse.step:type_name -> step.Step
	0,  // 4: step.ReadStepResponse.step:type_name -> step.Step
	2,  // 5: step.UpdateStepRequest.stepParams:type_name -> step.EditableStepParams
	0,  // 6: step.ListStepsResponse.steps:type_name -> step.Step
	3,  // 7: step.StepService.CreateStep:input_type -> step.CreateStepRequest
	5,  // 8: step.StepService.ReadStep:input_type -> step.ReadStepRequest
	7,  // 9: step.StepService.UpdateStep:input_type -> step.UpdateStepRequest
	9,  // 10: step.StepService.DeleteStep:input_type -> step.DeleteStepRequest
	11, // 11: step.StepService.ListSteps:input_type -> step.ListStepsRequest
	4,  // 12: step.StepService.CreateStep:output_type -> step.CreateStepResponse
	6,  // 13: step.StepService.ReadStep:output_type -> step.ReadStepResponse
	8,  // 14: step.StepService.UpdateStep:output_type -> step.UpdateStepResponse
	10, // 15: step.StepService.DeleteStep:output_type -> step.DeleteStepResponse
	12, // 16: step.StepService.ListSteps:output_type -> step.ListStepsResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_app_Adapter_In_ApiGrcp_proto_step_proto_init() }
func file_app_Adapter_In_ApiGrcp_proto_step_proto_init() {
	if File_app_Adapter_In_ApiGrcp_proto_step_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStepParams); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditableStepParams); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStepRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStepResponse); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadStepRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadStepResponse); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStepRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStepResponse); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStepRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStepResponse); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStepsRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStepsResponse); i {
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
			RawDescriptor: file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_Adapter_In_ApiGrcp_proto_step_proto_goTypes,
		DependencyIndexes: file_app_Adapter_In_ApiGrcp_proto_step_proto_depIdxs,
		MessageInfos:      file_app_Adapter_In_ApiGrcp_proto_step_proto_msgTypes,
	}.Build()
	File_app_Adapter_In_ApiGrcp_proto_step_proto = out.File
	file_app_Adapter_In_ApiGrcp_proto_step_proto_rawDesc = nil
	file_app_Adapter_In_ApiGrcp_proto_step_proto_goTypes = nil
	file_app_Adapter_In_ApiGrcp_proto_step_proto_depIdxs = nil
}
