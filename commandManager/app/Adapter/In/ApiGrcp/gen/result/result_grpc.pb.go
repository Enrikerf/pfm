// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: app/Adapter/In/ApiGrcp/proto/result.proto

package result

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

// ResultServiceClient is the client API for ResultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResultServiceClient interface {
	CreateResult(ctx context.Context, in *CreateResultRequest, opts ...grpc.CallOption) (*CreateResultResponse, error)
	ShowResult(ctx context.Context, in *ShowResultRequest, opts ...grpc.CallOption) (*ShowResultResponse, error)
	UpdateResult(ctx context.Context, in *UpdateResultRequest, opts ...grpc.CallOption) (*UpdateResultResponse, error)
	DeleteResult(ctx context.Context, in *DeleteResultRequest, opts ...grpc.CallOption) (*DeleteResultResponse, error)
	ListResult(ctx context.Context, in *ListResultRequest, opts ...grpc.CallOption) (*ListResultResponse, error)
}

type resultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResultServiceClient(cc grpc.ClientConnInterface) ResultServiceClient {
	return &resultServiceClient{cc}
}

func (c *resultServiceClient) CreateResult(ctx context.Context, in *CreateResultRequest, opts ...grpc.CallOption) (*CreateResultResponse, error) {
	out := new(CreateResultResponse)
	err := c.cc.Invoke(ctx, "/result.ResultService/CreateResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultServiceClient) ShowResult(ctx context.Context, in *ShowResultRequest, opts ...grpc.CallOption) (*ShowResultResponse, error) {
	out := new(ShowResultResponse)
	err := c.cc.Invoke(ctx, "/result.ResultService/ShowResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultServiceClient) UpdateResult(ctx context.Context, in *UpdateResultRequest, opts ...grpc.CallOption) (*UpdateResultResponse, error) {
	out := new(UpdateResultResponse)
	err := c.cc.Invoke(ctx, "/result.ResultService/UpdateResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultServiceClient) DeleteResult(ctx context.Context, in *DeleteResultRequest, opts ...grpc.CallOption) (*DeleteResultResponse, error) {
	out := new(DeleteResultResponse)
	err := c.cc.Invoke(ctx, "/result.ResultService/DeleteResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultServiceClient) ListResult(ctx context.Context, in *ListResultRequest, opts ...grpc.CallOption) (*ListResultResponse, error) {
	out := new(ListResultResponse)
	err := c.cc.Invoke(ctx, "/result.ResultService/ListResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultServiceServer is the server API for ResultService service.
// All implementations must embed UnimplementedResultServiceServer
// for forward compatibility
type ResultServiceServer interface {
	CreateResult(context.Context, *CreateResultRequest) (*CreateResultResponse, error)
	ShowResult(context.Context, *ShowResultRequest) (*ShowResultResponse, error)
	UpdateResult(context.Context, *UpdateResultRequest) (*UpdateResultResponse, error)
	DeleteResult(context.Context, *DeleteResultRequest) (*DeleteResultResponse, error)
	ListResult(context.Context, *ListResultRequest) (*ListResultResponse, error)
	mustEmbedUnimplementedResultServiceServer()
}

// UnimplementedResultServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResultServiceServer struct {
}

func (UnimplementedResultServiceServer) CreateResult(context.Context, *CreateResultRequest) (*CreateResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResult not implemented")
}
func (UnimplementedResultServiceServer) ShowResult(context.Context, *ShowResultRequest) (*ShowResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowResult not implemented")
}
func (UnimplementedResultServiceServer) UpdateResult(context.Context, *UpdateResultRequest) (*UpdateResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResult not implemented")
}
func (UnimplementedResultServiceServer) DeleteResult(context.Context, *DeleteResultRequest) (*DeleteResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResult not implemented")
}
func (UnimplementedResultServiceServer) ListResult(context.Context, *ListResultRequest) (*ListResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResult not implemented")
}
func (UnimplementedResultServiceServer) mustEmbedUnimplementedResultServiceServer() {}

// UnsafeResultServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResultServiceServer will
// result in compilation errors.
type UnsafeResultServiceServer interface {
	mustEmbedUnimplementedResultServiceServer()
}

func RegisterResultServiceServer(s grpc.ServiceRegistrar, srv ResultServiceServer) {
	s.RegisterService(&ResultService_ServiceDesc, srv)
}

func _ResultService_CreateResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).CreateResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultService/CreateResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).CreateResult(ctx, req.(*CreateResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultService_ShowResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).ShowResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultService/ShowResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).ShowResult(ctx, req.(*ShowResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultService_UpdateResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).UpdateResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultService/UpdateResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).UpdateResult(ctx, req.(*UpdateResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultService_DeleteResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).DeleteResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultService/DeleteResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).DeleteResult(ctx, req.(*DeleteResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultService_ListResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).ListResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultService/ListResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).ListResult(ctx, req.(*ListResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResultService_ServiceDesc is the grpc.ServiceDesc for ResultService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResultService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "result.ResultService",
	HandlerType: (*ResultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResult",
			Handler:    _ResultService_CreateResult_Handler,
		},
		{
			MethodName: "ShowResult",
			Handler:    _ResultService_ShowResult_Handler,
		},
		{
			MethodName: "UpdateResult",
			Handler:    _ResultService_UpdateResult_Handler,
		},
		{
			MethodName: "DeleteResult",
			Handler:    _ResultService_DeleteResult_Handler,
		},
		{
			MethodName: "ListResult",
			Handler:    _ResultService_ListResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/Adapter/In/ApiGrcp/proto/result.proto",
}
