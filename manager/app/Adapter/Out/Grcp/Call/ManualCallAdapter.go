package Call

import (
	"context"
	"errors"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"google.golang.org/grpc"
	"io"
	"time"
)

type ManualAdapter struct {
	connection *grpc.ClientConn
	client     call.CallService_CallBidirectionalClient
}

func (manualAdapter *ManualAdapter) Setup(host, port string) error {
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
func (manualAdapter *ManualAdapter) Write(step Entity.Step) error {
	callRequest := call.CallRequest{
		Step: step.Sentence,
	}
	err := manualAdapter.client.Send(&callRequest)
	if err != nil {
		return errors.New("send error")
	}
	return nil
}
func (manualAdapter *ManualAdapter) Read() (ValueObject.ResultVo, error) {
	response, err := manualAdapter.client.Recv()
	if err == io.EOF {
		return ValueObject.ResultVo{}, err
	}
	if err != nil {
		return ValueObject.ResultVo{}, err
	}
	return ValueObject.ResultVo{
		Content:   response.Result,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (manualAdapter *ManualAdapter) Close() {
	err := manualAdapter.connection.Close()
	if err != nil {
		return
	}
	err = manualAdapter.client.CloseSend()
	if err != nil {
		return
	}
}
