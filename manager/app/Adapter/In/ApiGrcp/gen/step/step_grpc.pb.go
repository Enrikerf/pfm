// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: app/Adapter/In/ApiGrcp/proto/step.proto

package step

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

// StepServiceClient is the client API for StepService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StepServiceClient interface {
	CreateStep(ctx context.Context, in *CreateStepRequest, opts ...grpc.CallOption) (*CreateStepResponse, error)
	ReadStep(ctx context.Context, in *ReadStepRequest, opts ...grpc.CallOption) (*ReadStepResponse, error)
	UpdateStep(ctx context.Context, in *UpdateStepRequest, opts ...grpc.CallOption) (*UpdateStepResponse, error)
	DeleteStep(ctx context.Context, in *DeleteStepRequest, opts ...grpc.CallOption) (*DeleteStepResponse, error)
	ListSteps(ctx context.Context, in *ListStepsRequest, opts ...grpc.CallOption) (*ListStepsResponse, error)
}

type stepServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStepServiceClient(cc grpc.ClientConnInterface) StepServiceClient {
	return &stepServiceClient{cc}
}

func (c *stepServiceClient) CreateStep(ctx context.Context, in *CreateStepRequest, opts ...grpc.CallOption) (*CreateStepResponse, error) {
	out := new(CreateStepResponse)
	err := c.cc.Invoke(ctx, "/step.StepService/CreateStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stepServiceClient) ReadStep(ctx context.Context, in *ReadStepRequest, opts ...grpc.CallOption) (*ReadStepResponse, error) {
	out := new(ReadStepResponse)
	err := c.cc.Invoke(ctx, "/step.StepService/ReadStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stepServiceClient) UpdateStep(ctx context.Context, in *UpdateStepRequest, opts ...grpc.CallOption) (*UpdateStepResponse, error) {
	out := new(UpdateStepResponse)
	err := c.cc.Invoke(ctx, "/step.StepService/UpdateStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stepServiceClient) DeleteStep(ctx context.Context, in *DeleteStepRequest, opts ...grpc.CallOption) (*DeleteStepResponse, error) {
	out := new(DeleteStepResponse)
	err := c.cc.Invoke(ctx, "/step.StepService/DeleteStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stepServiceClient) ListSteps(ctx context.Context, in *ListStepsRequest, opts ...grpc.CallOption) (*ListStepsResponse, error) {
	out := new(ListStepsResponse)
	err := c.cc.Invoke(ctx, "/step.StepService/ListSteps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StepServiceServer is the server API for StepService service.
// All implementations must embed UnimplementedStepServiceServer
// for forward compatibility
type StepServiceServer interface {
	CreateStep(context.Context, *CreateStepRequest) (*CreateStepResponse, error)
	ReadStep(context.Context, *ReadStepRequest) (*ReadStepResponse, error)
	UpdateStep(context.Context, *UpdateStepRequest) (*UpdateStepResponse, error)
	DeleteStep(context.Context, *DeleteStepRequest) (*DeleteStepResponse, error)
	ListSteps(context.Context, *ListStepsRequest) (*ListStepsResponse, error)
	mustEmbedUnimplementedStepServiceServer()
}

// UnimplementedStepServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStepServiceServer struct {
}

func (UnimplementedStepServiceServer) CreateStep(context.Context, *CreateStepRequest) (*CreateStepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStep not implemented")
}
func (UnimplementedStepServiceServer) ReadStep(context.Context, *ReadStepRequest) (*ReadStepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadStep not implemented")
}
func (UnimplementedStepServiceServer) UpdateStep(context.Context, *UpdateStepRequest) (*UpdateStepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStep not implemented")
}
func (UnimplementedStepServiceServer) DeleteStep(context.Context, *DeleteStepRequest) (*DeleteStepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStep not implemented")
}
func (UnimplementedStepServiceServer) ListSteps(context.Context, *ListStepsRequest) (*ListStepsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSteps not implemented")
}
func (UnimplementedStepServiceServer) mustEmbedUnimplementedStepServiceServer() {}

// UnsafeStepServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StepServiceServer will
// result in compilation errors.
type UnsafeStepServiceServer interface {
	mustEmbedUnimplementedStepServiceServer()
}

func RegisterStepServiceServer(s grpc.ServiceRegistrar, srv StepServiceServer) {
	s.RegisterService(&StepService_ServiceDesc, srv)
}

func _StepService_CreateStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StepServiceServer).CreateStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/step.StepService/CreateStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StepServiceServer).CreateStep(ctx, req.(*CreateStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StepService_ReadStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StepServiceServer).ReadStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/step.StepService/ReadStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StepServiceServer).ReadStep(ctx, req.(*ReadStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StepService_UpdateStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StepServiceServer).UpdateStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/step.StepService/UpdateStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StepServiceServer).UpdateStep(ctx, req.(*UpdateStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StepService_DeleteStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StepServiceServer).DeleteStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/step.StepService/DeleteStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StepServiceServer).DeleteStep(ctx, req.(*DeleteStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StepService_ListSteps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StepServiceServer).ListSteps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/step.StepService/ListSteps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StepServiceServer).ListSteps(ctx, req.(*ListStepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StepService_ServiceDesc is the grpc.ServiceDesc for StepService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StepService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "step.StepService",
	HandlerType: (*StepServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStep",
			Handler:    _StepService_CreateStep_Handler,
		},
		{
			MethodName: "ReadStep",
			Handler:    _StepService_ReadStep_Handler,
		},
		{
			MethodName: "UpdateStep",
			Handler:    _StepService_UpdateStep_Handler,
		},
		{
			MethodName: "DeleteStep",
			Handler:    _StepService_DeleteStep_Handler,
		},
		{
			MethodName: "ListSteps",
			Handler:    _StepService_ListSteps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/Adapter/In/ApiGrcp/proto/step.proto",
}
