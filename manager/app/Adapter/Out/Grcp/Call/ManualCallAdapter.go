package Call

import (
	"context"
	"errors"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"google.golang.org/grpc"
	"io"
)

type ManualAdapter interface {
	Setup(host, port string) error
	Write(step Step.Step) error
	Read() (Content.Content, error)
	Close()
}

func New() ManualAdapter {
	return &manualAdapter{}
}

type manualAdapter struct {
	connection *grpc.ClientConn
	client     call.CallService_CallBidirectionalClient
}

func (manualAdapter *manualAdapter) Setup(host, port string) error {
	options := grpc.WithInsecure()
	var err error
	manualAdapter.connection, err = grpc.Dial(host+":"+port, options)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return errors.New("can't setup")
	}
	client := call.NewCallServiceClient(manualAdapter.connection)
	manualAdapter.client, err = client.CallBidirectional(context.Background())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return errors.New("can't setup")
	}
	return nil
}
func (manualAdapter *manualAdapter) Write(step Step.Step) error {
	callRequest := call.CallRequest{
		Step: step.GetSentence(),
	}
	err := manualAdapter.client.Send(&callRequest)
	if err != nil {
		return errors.New("send error")
	}
	return nil
}
func (manualAdapter *manualAdapter) Read() (Content.Content, error) {
	response, err := manualAdapter.client.Recv()
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return Content.NewContent(response.Result)
}
func (manualAdapter *manualAdapter) Close() {
	err := manualAdapter.connection.Close()
	if err != nil {
		return
	}
	err = manualAdapter.client.CloseSend()
	if err != nil {
		return
	}
}
