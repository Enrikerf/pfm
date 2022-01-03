package Call

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"google.golang.org/grpc"
)

type Adapter struct {
}

func (adapter Adapter) Request(task Task.Task) Result.Result {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.Host+":"+task.Port, options)
	if err != nil {
		fmt.Println("error: %v", err)
	}
	defer connection.Close()
	client := call.NewCallServiceClient(connection)
	callRequest := call.CallRequest{
		Command: task.Command,
	}
	result, _ := Result.NewResult(task.Id,"")
	switch task.Mode {
	case Task.Unary:
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
			result.Content = err.Error()
		}
		result.Content = callResponse.GetResult()
	case Task.ServerStream:
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
			result.Content = err.Error()
		}
		result.Content = callResponse.GetResult()
	case Task.ClientStream:
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
			result.Content = err.Error()
		}
		result.Content = callResponse.GetResult()
	case Task.Bidirectional:
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
			result.Content = err.Error()
		}
		result.Content = callResponse.GetResult()
	default:
		result.Content = "not implemented task mode"
		fmt.Printf("%v) task:  \n", err)
	}

	return result
}
