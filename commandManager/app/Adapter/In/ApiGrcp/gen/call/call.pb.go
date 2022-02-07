// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: app/Adapter/In/ApiGrcp/proto/call.proto

package call

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

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step string `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetStep() string {
	if x != nil {
		return x.Step
	}
	return ""
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_app_Adapter_In_ApiGrcp_proto_call_proto protoreflect.FileDescriptor

var file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x70, 0x2f, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x49, 0x6e,
	0x2f, 0x41, 0x70, 0x69, 0x47, 0x72, 0x63, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x22,
	0x21, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x22, 0x26, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x83, 0x02, 0x0a, 0x0b, 0x43,
	0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x61,
	0x6c, 0x6c, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c,
	0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x3d, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x40,
	0x0a, 0x11, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescOnce sync.Once
	file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescData = file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDesc
)

func file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescGZIP() []byte {
	file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescOnce.Do(func() {
		file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescData)
	})
	return file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDescData
}

var file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_Adapter_In_ApiGrcp_proto_call_proto_goTypes = []interface{}{
	(*CallRequest)(nil),  // 0: call.CallRequest
	(*CallResponse)(nil), // 1: call.CallResponse
}
var file_app_Adapter_In_ApiGrcp_proto_call_proto_depIdxs = []int32{
	0, // 0: call.CallService.CallUnary:input_type -> call.CallRequest
	0, // 1: call.CallService.CallServerStream:input_type -> call.CallRequest
	0, // 2: call.CallService.CallClientStream:input_type -> call.CallRequest
	0, // 3: call.CallService.CallBidirectional:input_type -> call.CallRequest
	1, // 4: call.CallService.CallUnary:output_type -> call.CallResponse
	1, // 5: call.CallService.CallServerStream:output_type -> call.CallResponse
	1, // 6: call.CallService.CallClientStream:output_type -> call.CallResponse
	1, // 7: call.CallService.CallBidirectional:output_type -> call.CallResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_Adapter_In_ApiGrcp_proto_call_proto_init() }
func file_app_Adapter_In_ApiGrcp_proto_call_proto_init() {
	if File_app_Adapter_In_ApiGrcp_proto_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
			RawDescriptor: file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_Adapter_In_ApiGrcp_proto_call_proto_goTypes,
		DependencyIndexes: file_app_Adapter_In_ApiGrcp_proto_call_proto_depIdxs,
		MessageInfos:      file_app_Adapter_In_ApiGrcp_proto_call_proto_msgTypes,
	}.Build()
	File_app_Adapter_In_ApiGrcp_proto_call_proto = out.File
	file_app_Adapter_In_ApiGrcp_proto_call_proto_rawDesc = nil
	file_app_Adapter_In_ApiGrcp_proto_call_proto_goTypes = nil
	file_app_Adapter_In_ApiGrcp_proto_call_proto_depIdxs = nil
}
