// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: app/Adapter/In/ApiGrcp/proto/batch.proto

package batch

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BatchServiceClient is the client API for BatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BatchServiceClient interface {
	CreateBatch(ctx context.Context, in *CreateBatchRequest, opts ...grpc.CallOption) (*CreateBatchResponse, error)
	ReadBatch(ctx context.Context, in *ReadBatchRequest, opts ...grpc.CallOption) (*ReadBatchResponse, error)
	UpdateBatch(ctx context.Context, in *UpdateBatchRequest, opts ...grpc.CallOption) (*UpdateBatchResponse, error)
	DeleteBatch(ctx context.Context, in *DeleteBatchRequest, opts ...grpc.CallOption) (*DeleteBatchResponse, error)
	ListBatches(ctx context.Context, in *ListBatchesRequest, opts ...grpc.CallOption) (*ListBatchesResponse, error)
}

type batchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBatchServiceClient(cc grpc.ClientConnInterface) BatchServiceClient {
	return &batchServiceClient{cc}
}

func (c *batchServiceClient) CreateBatch(ctx context.Context, in *CreateBatchRequest, opts ...grpc.CallOption) (*CreateBatchResponse, error) {
	out := new(CreateBatchResponse)
	err := c.cc.Invoke(ctx, "/batch.BatchService/CreateBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchServiceClient) ReadBatch(ctx context.Context, in *ReadBatchRequest, opts ...grpc.CallOption) (*ReadBatchResponse, error) {
	out := new(ReadBatchResponse)
	err := c.cc.Invoke(ctx, "/batch.BatchService/ReadBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchServiceClient) UpdateBatch(ctx context.Context, in *UpdateBatchRequest, opts ...grpc.CallOption) (*UpdateBatchResponse, error) {
	out := new(UpdateBatchResponse)
	err := c.cc.Invoke(ctx, "/batch.BatchService/UpdateBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchServiceClient) DeleteBatch(ctx context.Context, in *DeleteBatchRequest, opts ...grpc.CallOption) (*DeleteBatchResponse, error) {
	out := new(DeleteBatchResponse)
	err := c.cc.Invoke(ctx, "/batch.BatchService/DeleteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchServiceClient) ListBatches(ctx context.Context, in *ListBatchesRequest, opts ...grpc.CallOption) (*ListBatchesResponse, error) {
	out := new(ListBatchesResponse)
	err := c.cc.Invoke(ctx, "/batch.BatchService/ListBatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchServiceServer is the server API for BatchService service.
// All implementations must embed UnimplementedBatchServiceServer
// for forward compatibility
type BatchServiceServer interface {
	CreateBatch(context.Context, *CreateBatchRequest) (*CreateBatchResponse, error)
	ReadBatch(context.Context, *ReadBatchRequest) (*ReadBatchResponse, error)
	UpdateBatch(context.Context, *UpdateBatchRequest) (*UpdateBatchResponse, error)
	DeleteBatch(context.Context, *DeleteBatchRequest) (*DeleteBatchResponse, error)
	ListBatches(context.Context, *ListBatchesRequest) (*ListBatchesResponse, error)
	mustEmbedUnimplementedBatchServiceServer()
}

// UnimplementedBatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBatchServiceServer struct {
}

func (UnimplementedBatchServiceServer) CreateBatch(context.Context, *CreateBatchRequest) (*CreateBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatch not implemented")
}
func (UnimplementedBatchServiceServer) ReadBatch(context.Context, *ReadBatchRequest) (*ReadBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBatch not implemented")
}
func (UnimplementedBatchServiceServer) UpdateBatch(context.Context, *UpdateBatchRequest) (*UpdateBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBatch not implemented")
}
func (UnimplementedBatchServiceServer) DeleteBatch(context.Context, *DeleteBatchRequest) (*DeleteBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBatch not implemented")
}
func (UnimplementedBatchServiceServer) ListBatches(context.Context, *ListBatchesRequest) (*ListBatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBatches not implemented")
}
func (UnimplementedBatchServiceServer) mustEmbedUnimplementedBatchServiceServer() {}

// UnsafeBatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatchServiceServer will
// result in compilation errors.
type UnsafeBatchServiceServer interface {
	mustEmbedUnimplementedBatchServiceServer()
}

func RegisterBatchServiceServer(s grpc.ServiceRegistrar, srv BatchServiceServer) {
	s.RegisterService(&BatchService_ServiceDesc, srv)
}

func _BatchService_CreateBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).CreateBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch.BatchService/CreateBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).CreateBatch(ctx, req.(*CreateBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchService_ReadBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).ReadBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch.BatchService/ReadBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).ReadBatch(ctx, req.(*ReadBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchService_UpdateBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).UpdateBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch.BatchService/UpdateBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).UpdateBatch(ctx, req.(*UpdateBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchService_DeleteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).DeleteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch.BatchService/DeleteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).DeleteBatch(ctx, req.(*DeleteBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchService_ListBatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchServiceServer).ListBatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/batch.BatchService/ListBatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchServiceServer).ListBatches(ctx, req.(*ListBatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BatchService_ServiceDesc is the grpc.ServiceDesc for BatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "batch.BatchService",
	HandlerType: (*BatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBatch",
			Handler:    _BatchService_CreateBatch_Handler,
		},
		{
			MethodName: "ReadBatch",
			Handler:    _BatchService_ReadBatch_Handler,
		},
		{
			MethodName: "UpdateBatch",
			Handler:    _BatchService_UpdateBatch_Handler,
		},
		{
			MethodName: "DeleteBatch",
			Handler:    _BatchService_DeleteBatch_Handler,
		},
		{
			MethodName: "ListBatches",
			Handler:    _BatchService_ListBatches_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/Adapter/In/ApiGrcp/proto/batch.proto",
}
