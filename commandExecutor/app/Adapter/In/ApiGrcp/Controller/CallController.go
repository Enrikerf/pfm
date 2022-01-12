package Controller

import (
	"context"
	"github.com/Enrikerf/pfm/commandExecutor/app/Adapter/In/ApiGrcp/gen/call"
)

type CallController struct {
	call.UnimplementedCallServiceServer
}

func (s CallController) CallUnary(ctx context.Context, request *call.CallRequest) (*call.CallResponse, error) {
	println("unary")
	return &call.CallResponse{Result: "result"}, nil
}

func (s CallController) CallServerStream(request *call.CallRequest, server call.CallService_CallServerStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (s CallController) CallClientStream(server call.CallService_CallClientStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (s CallController) CallBidirectional(server call.CallService_CallBidirectionalServer) error {
	//TODO implement me
	panic("implement me")
}
