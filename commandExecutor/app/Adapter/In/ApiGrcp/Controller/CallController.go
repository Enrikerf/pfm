package Controller

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandExecutor/app/Adapter/In/ApiGrcp/gen/call"
	"os/exec"
	"strings"
)

type CallController struct {
	call.UnimplementedCallServiceServer
}

func (s CallController) CallUnary(ctx context.Context, request *call.CallRequest) (*call.CallResponse, error) {
	var resultContent string
	parts := strings.Fields(request.Command)
	head := parts[0]
	// Head at this point is "sudo"
	parts = parts[1:len(parts)]
	cmd := exec.Command(head, parts...)
	stdout, err := cmd.Output()
	if err != nil {
		resultContent = err.Error()
		fmt.Println(err.Error())
	} else {
		resultContent = string(stdout)
	}

	return &call.CallResponse{Result: resultContent}, nil
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
