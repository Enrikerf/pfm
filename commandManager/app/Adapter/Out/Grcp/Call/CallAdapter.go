package Call

import (
	"context"
	"fmt"
	ProtoCall "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	DomainResult "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	DomainTask "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"google.golang.org/grpc"
	"io"
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
		results = adapter.doUnaryCall(task, client)
	case DomainTask.ServerStream:
		results = adapter.notImplementedMethod(task)
	case DomainTask.ClientStream:
		results = adapter.doClientStream(task, client)
	case DomainTask.Bidirectional:
		results = adapter.doBidirectional(task, client)
	default:
		results = adapter.notImplementedMethod(task)
	}

	return results
}

func (adapter Adapter) doBidirectional(task DomainTask.Task, client ProtoCall.CallServiceClient) []DomainResult.Result {
	results := []DomainResult.Result{}
	fmt.Println("find maximum")
	stream, err := client.CallBidirectional(context.Background())
	if err != nil {
		result, _ := DomainResult.NewResult(task.Uuid, "")
		result.Content = err.Error()
		results = append(results, result)
	}
	// we send a bunch of messages to de client go routine
	resultsChannel1 := make(chan string)
	go sendInStream(task, stream, resultsChannel1)
	// we receive a bunch of messages form the client go routine
	resultsChannel2 := make(chan string)
	go receiveServerStream(stream, resultsChannel2)
	for i := range resultsChannel1 {
		result, _ := DomainResult.NewResult(task.Uuid, "")
		result.Content = i
		results = append(results, result)
	}
	for i := range resultsChannel2 {
		result, _ := DomainResult.NewResult(task.Uuid, "")
		result.Content = i
		results = append(results, result)
	}
	return results
}

func receiveServerStream(stream ProtoCall.CallService_CallBidirectionalClient, resultsChannel chan string) {
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			resultsChannel <- err.Error()
			break
		}
		if err != nil {
			resultsChannel <- err.Error()
			break
		}
		resultsChannel <- response.Result
		fmt.Println("received: %v", response)
	}
	close(resultsChannel)
}

func sendInStream(task DomainTask.Task, stream ProtoCall.CallService_CallBidirectionalClient, resultsChannel chan string) {
	for _, command := range task.Commands {
		fmt.Println("sending message %v", command)
		callRequest := ProtoCall.CallRequest{
			Command: command.Name,
		}
		err := stream.Send(&callRequest)
		if err != nil {
			resultsChannel <- err.Error()
		}
	}
	err := stream.CloseSend()
	if err != nil {
		resultsChannel <- err.Error()
	}
	close(resultsChannel)
}

func (adapter Adapter) doClientStream(task DomainTask.Task, client ProtoCall.CallServiceClient) []DomainResult.Result {
	results := []DomainResult.Result{}
	result, _ := DomainResult.NewResult(task.Uuid, "")
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
				result.Content = err.Error()
				results = append(results, result)
			}
		}
		response, err := stream.CloseAndRecv()
		if err != nil {
			result.Content = err.Error()
			results = append(results, result)
		}
		result, _ := DomainResult.NewResult(task.Uuid, response.GetResult())
		results = append(results, result)
	}
	return results
}

func (adapter Adapter) doUnaryCall(task DomainTask.Task, client ProtoCall.CallServiceClient) []DomainResult.Result {
	results := []DomainResult.Result{}
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

func (adapter Adapter) notImplementedMethod(task DomainTask.Task) []DomainResult.Result {
	results := []DomainResult.Result{}
	result, _ := DomainResult.NewResult(task.Uuid, "not implemented")
	results = append(results, result)
	return results
}
