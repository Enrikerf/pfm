package Call

import (
	"context"
	"fmt"
	ProtoCall "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	DomainResult "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	DomainTask "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"google.golang.org/grpc"
)

type Adapter struct {
}

func (adapter Adapter) Request(task DomainTask.Task) []DomainResult.Result {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.Host+":"+task.Port, options)
	if err != nil {
		fmt.Println("error: %v", err)
	}
	defer connection.Close()
	client := ProtoCall.NewCallServiceClient(connection)
	results := []DomainResult.Result{}
	switch task.Mode {
	case DomainTask.Unary:
		results = adapter.doUnaryCall(task, client, results)
	case DomainTask.ServerStream:
		results = adapter.notImplementedMethod(task, results)
	case DomainTask.ClientStream:
		result, _ := DomainResult.NewResult(task.Uuid, "not implemented")
		stream, err := client.CallClientStream(context.Background())
		if err != nil {
			result.Content = err.Error()
		} else {
			for _, command := range task.Commands {
				fmt.Printf("sending: %v\n", command.Name)
				callRequest := ProtoCall.CallRequest{
					Command: command.Name,
				}
				err := stream.Send(&callRequest)
				if err != nil {
					return nil
				}
			}
			response, error := stream.CloseAndRecv()
			if error != nil {
				return nil
			}
			result, _ := DomainResult.NewResult(task.Uuid, response.GetResult())
			results = append(results, result)

		}
	case DomainTask.Bidirectional:
		results = adapter.notImplementedMethod(task, results)
	default:
		results = adapter.notImplementedMethod(task, results)
	}

	return results
}

func (adapter Adapter) doUnaryCall(task DomainTask.Task, client ProtoCall.CallServiceClient, results []DomainResult.Result) []DomainResult.Result {
	callRequest := ProtoCall.CallRequest{
		Command: task.Commands[0].Name,
	}
	callResponse, err := client.CallUnary(context.Background(), &callRequest)
	result, _ := DomainResult.NewResult(task.Uuid, "")
	if err != nil {
		result.Content = err.Error()
	} else {
		result.Content = callResponse.GetResult()
	}
	results = append(results, result)
	return results
}

func (adapter Adapter) notImplementedMethod(task DomainTask.Task, results []DomainResult.Result) []DomainResult.Result {
	result, _ := DomainResult.NewResult(task.Uuid, "not implemented")
	results = append(results, result)
	return results
}
