// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: app/Adapter/In/ApiGrcp/proto/result.proto

package result

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

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TaskUuid  string `protobuf:"bytes,2,opt,name=task_uuid,json=taskUuid,proto3" json:"task_uuid,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{0}
}

func (x *Batch) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Batch) GetTaskUuid() string {
	if x != nil {
		return x.TaskUuid
	}
	return ""
}

func (x *Batch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CommunicateTaskManuallyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskUuid string `protobuf:"bytes,1,opt,name=task_uuid,json=taskUuid,proto3" json:"task_uuid,omitempty"`
}

func (x *CommunicateTaskManuallyRequest) Reset() {
	*x = CommunicateTaskManuallyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicateTaskManuallyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicateTaskManuallyRequest) ProtoMessage() {}

func (x *CommunicateTaskManuallyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicateTaskManuallyRequest.ProtoReflect.Descriptor instead.
func (*CommunicateTaskManuallyRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{1}
}

func (x *CommunicateTaskManuallyRequest) GetTaskUuid() string {
	if x != nil {
		return x.TaskUuid
	}
	return ""
}

type CommunicateTaskManuallyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchUuid string `protobuf:"bytes,1,opt,name=batch_uuid,json=batchUuid,proto3" json:"batch_uuid,omitempty"`
}

func (x *CommunicateTaskManuallyResponse) Reset() {
	*x = CommunicateTaskManuallyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicateTaskManuallyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicateTaskManuallyResponse) ProtoMessage() {}

func (x *CommunicateTaskManuallyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicateTaskManuallyResponse.ProtoReflect.Descriptor instead.
func (*CommunicateTaskManuallyResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{2}
}

func (x *CommunicateTaskManuallyResponse) GetBatchUuid() string {
	if x != nil {
		return x.BatchUuid
	}
	return ""
}

type GetTaskBatchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskUuid string `protobuf:"bytes,1,opt,name=task_uuid,json=taskUuid,proto3" json:"task_uuid,omitempty"`
}

func (x *GetTaskBatchesRequest) Reset() {
	*x = GetTaskBatchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskBatchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskBatchesRequest) ProtoMessage() {}

func (x *GetTaskBatchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskBatchesRequest.ProtoReflect.Descriptor instead.
func (*GetTaskBatchesRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskBatchesRequest) GetTaskUuid() string {
	if x != nil {
		return x.TaskUuid
	}
	return ""
}

type ListBatchesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batches []*Batch `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
}

func (x *ListBatchesResponse) Reset() {
	*x = ListBatchesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchesResponse) ProtoMessage() {}

func (x *ListBatchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchesResponse.ProtoReflect.Descriptor instead.
func (*ListBatchesResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{4}
}

func (x *ListBatchesResponse) GetBatches() []*Batch {
	if x != nil {
		return x.Batches
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	BatchUuid string `protobuf:"bytes,2,opt,name=batch_uuid,json=batchUuid,proto3" json:"batch_uuid,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{5}
}

func (x *Result) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Result) GetBatchUuid() string {
	if x != nil {
		return x.BatchUuid
	}
	return ""
}

func (x *Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Result) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchUuid *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=batch_uuid,json=batchUuid,proto3" json:"batch_uuid,omitempty"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{6}
}

func (x *Filters) GetBatchUuid() *wrapperspb.StringValue {
	if x != nil {
		return x.BatchUuid
	}
	return nil
}

type GetBatchResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchUuid string `protobuf:"bytes,1,opt,name=batch_uuid,json=batchUuid,proto3" json:"batch_uuid,omitempty"`
}

func (x *GetBatchResultsRequest) Reset() {
	*x = GetBatchResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchResultsRequest) ProtoMessage() {}

func (x *GetBatchResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchResultsRequest.ProtoReflect.Descriptor instead.
func (*GetBatchResultsRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{7}
}

func (x *GetBatchResultsRequest) GetBatchUuid() string {
	if x != nil {
		return x.BatchUuid
	}
	return ""
}

type ListResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListResultsResponse) Reset() {
	*x = ListResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResultsResponse) ProtoMessage() {}

func (x *ListResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResultsResponse.ProtoReflect.Descriptor instead.
func (*ListResultsResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{8}
}

func (x *ListResultsResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type StreamResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchUuid string `protobuf:"bytes,1,opt,name=batch_uuid,json=batchUuid,proto3" json:"batch_uuid,omitempty"`
}

func (x *StreamResultsRequest) Reset() {
	*x = StreamResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResultsRequest) ProtoMessage() {}

func (x *StreamResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResultsRequest.ProtoReflect.Descriptor instead.
func (*StreamResultsRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{9}
}

func (x *StreamResultsRequest) GetBatchUuid() string {
	if x != nil {
		return x.BatchUuid
	}
	return ""
}

type StreamResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *StreamResultsResponse) Reset() {
	*x = StreamResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResultsResponse) ProtoMessage() {}

func (x *StreamResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResultsResponse.ProtoReflect.Descriptor instead.
func (*StreamResultsResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP(), []int{10}
}

func (x *StreamResultsResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_app_Adapter_In_ApiGrcp_proto_result_proto protoreflect.FileDescriptor

var file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x70, 0x2f, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x49, 0x6e,
	0x2f, 0x41, 0x70, 0x69, 0x47, 0x72, 0x63, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3d, 0x0a, 0x1e,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x55, 0x75, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x1f, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x55, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x55,
	0x75, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x22, 0x74, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x46, 0x0a, 0x07, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x55, 0x75, 0x69,
	0x64, 0x22, 0x37, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x55, 0x75, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x55, 0x75,
	0x69, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0xe9, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x6c, 0x79, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescOnce sync.Once
	file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescData = file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDesc
)

func file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescGZIP() []byte {
	file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescOnce.Do(func() {
		file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescData)
	})
	return file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDescData
}

var file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_app_Adapter_In_ApiGrcp_proto_result_proto_goTypes = []interface{}{
	(*Batch)(nil),                           // 0: result.Batch
	(*CommunicateTaskManuallyRequest)(nil),  // 1: result.CommunicateTaskManuallyRequest
	(*CommunicateTaskManuallyResponse)(nil), // 2: result.CommunicateTaskManuallyResponse
	(*GetTaskBatchesRequest)(nil),           // 3: result.GetTaskBatchesRequest
	(*ListBatchesResponse)(nil),             // 4: result.ListBatchesResponse
	(*Result)(nil),                          // 5: result.Result
	(*Filters)(nil),                         // 6: result.Filters
	(*GetBatchResultsRequest)(nil),          // 7: result.GetBatchResultsRequest
	(*ListResultsResponse)(nil),             // 8: result.ListResultsResponse
	(*StreamResultsRequest)(nil),            // 9: result.StreamResultsRequest
	(*StreamResultsResponse)(nil),           // 10: result.StreamResultsResponse
	(*wrapperspb.StringValue)(nil),          // 11: google.protobuf.StringValue
}
var file_app_Adapter_In_ApiGrcp_proto_result_proto_depIdxs = []int32{
	0,  // 0: result.ListBatchesResponse.batches:type_name -> result.Batch
	11, // 1: result.Filters.batch_uuid:type_name -> google.protobuf.StringValue
	5,  // 2: result.ListResultsResponse.results:type_name -> result.Result
	5,  // 3: result.StreamResultsResponse.results:type_name -> result.Result
	1,  // 4: result.ResultService.CommunicateTaskManually:input_type -> result.CommunicateTaskManuallyRequest
	7,  // 5: result.ResultService.GetBatchResults:input_type -> result.GetBatchResultsRequest
	3,  // 6: result.ResultService.GetTaskBatches:input_type -> result.GetTaskBatchesRequest
	9,  // 7: result.ResultService.StreamResults:input_type -> result.StreamResultsRequest
	2,  // 8: result.ResultService.CommunicateTaskManually:output_type -> result.CommunicateTaskManuallyResponse
	8,  // 9: result.ResultService.GetBatchResults:output_type -> result.ListResultsResponse
	4,  // 10: result.ResultService.GetTaskBatches:output_type -> result.ListBatchesResponse
	10, // 11: result.ResultService.StreamResults:output_type -> result.StreamResultsResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_app_Adapter_In_ApiGrcp_proto_result_proto_init() }
func file_app_Adapter_In_ApiGrcp_proto_result_proto_init() {
	if File_app_Adapter_In_ApiGrcp_proto_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicateTaskManuallyRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicateTaskManuallyResponse); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskBatchesRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchesResponse); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filters); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchResultsRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResultsResponse); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResultsRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResultsResponse); i {
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
			RawDescriptor: file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_Adapter_In_ApiGrcp_proto_result_proto_goTypes,
		DependencyIndexes: file_app_Adapter_In_ApiGrcp_proto_result_proto_depIdxs,
		MessageInfos:      file_app_Adapter_In_ApiGrcp_proto_result_proto_msgTypes,
	}.Build()
	File_app_Adapter_In_ApiGrcp_proto_result_proto = out.File
	file_app_Adapter_In_ApiGrcp_proto_result_proto_rawDesc = nil
	file_app_Adapter_In_ApiGrcp_proto_result_proto_goTypes = nil
	file_app_Adapter_In_ApiGrcp_proto_result_proto_depIdxs = nil
}
