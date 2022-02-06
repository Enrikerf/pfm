package Call

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"google.golang.org/grpc"
	"io"
)

type Adapter struct {
}

func (adapter Adapter) Request(task Entity.Task) Entity.Batch {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.Host+":"+task.Port, options)
	results := []Entity.Result{}
	if err != nil {
		fmt.Println("error: %v", err)
		result, _ := Entity.NewResult(task.Uuid, err.Error())
		results = append(results, result)
	} else {
		defer connection.Close()
		client := call.NewCallServiceClient(connection)
		switch task.Mode {
		case ValueObject.Unary:
			results = adapter.doUnaryCall(task, client)
		case ValueObject.ServerStream:
			results = adapter.doServerStream(task, client)
		case ValueObject.ClientStream:
			results = adapter.doClientStream(task, client)
		case ValueObject.Bidirectional:
			results = adapter.doBidirectional(task, client)
		default:
			results = adapter.notImplementedMethod(task)
		}
	}

	return Entity.NewBatch(task.Uuid, results)
}

func (adapter Adapter) doServerStream(task Entity.Task, client call.CallServiceClient) []Entity.Result {
	results := []Entity.Result{}
	request := &call.CallRequest{
		Step: task.Steps[0].Sentence,
	}
	responseStream, err := client.CallServerStream(context.Background(), request)

	if err != nil {
		result, _ := Entity.NewResult(task.Uuid, err.Error())
		results = append(results, result)
	} else {
		for {
			msg, err := responseStream.Recv()
			if err == io.EOF {
				result, _ := Entity.NewResult(task.Uuid, err.Error())
				results = append(results, result)
				break
			}
			if err != nil {
				result, _ := Entity.NewResult(task.Uuid, err.Error())
				results = append(results, result)
				break
			}
			result, _ := Entity.NewResult(task.Uuid, msg.GetResult())
			results = append(results, result)
		}
	}
	return results
}

func (adapter Adapter) doBidirectional(task Entity.Task, client call.CallServiceClient) []Entity.Result {
	results := []Entity.Result{}
	fmt.Println("find maximum")
	stream, err := client.CallBidirectional(context.Background())
	if err != nil {
		result, _ := Entity.NewResult(task.Uuid, "")
		result.Content = err.Error()
		results = append(results, result)
	} else {
		// we send a bunch of messages to de client go routine
		resultsChannel1 := make(chan string)
		go sendInStream(task, stream, resultsChannel1)
		// we receive a bunch of messages form the client go routine
		resultsChannel2 := make(chan string)
		go receiveServerStream(stream, resultsChannel2)
		for i := range resultsChannel1 {
			result, _ := Entity.NewResult(task.Uuid, "")
			result.Content = i
			results = append(results, result)
		}
		for i := range resultsChannel2 {
			result, _ := Entity.NewResult(task.Uuid, "")
			result.Content = i
			results = append(results, result)
		}
	}

	return results
}

func receiveServerStream(stream call.CallService_CallBidirectionalClient, resultsChannel chan string) {
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

func sendInStream(task Entity.Task, stream call.CallService_CallBidirectionalClient, resultsChannel chan string) {
	for _, step := range task.Steps {
		fmt.Println("sending message %v", step)
		callRequest := call.CallRequest{
			Step: step.Sentence,
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

func (adapter Adapter) doClientStream(task Entity.Task, client call.CallServiceClient) []Entity.Result {
	results := []Entity.Result{}
	result, _ := Entity.NewResult(task.Uuid, "")
	stream, err := client.CallClientStream(context.Background())
	if err != nil {
		result.Content = err.Error()
	} else {
		for _, step := range task.Steps {
			fmt.Printf("sending: %v\n", step.Sentence)
			callRequest := call.CallRequest{
				Step: step.Sentence,
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
		result, _ := Entity.NewResult(task.Uuid, response.GetResult())
		results = append(results, result)
	}
	return results
}

func (adapter Adapter) doUnaryCall(task Entity.Task, client call.CallServiceClient) []Entity.Result {
	results := []Entity.Result{}
	callRequest := call.CallRequest{
		Step: task.Steps[0].Sentence,
	}
	callResponse, err := client.CallUnary(context.Background(), &callRequest)
	result, _ := Entity.NewResult(task.Uuid, "")
	if err != nil {
		result.Content = err.Error()
	} else {
		result.Content = callResponse.GetResult()
	}
	results = append(results, result)
	return results
}

func (adapter Adapter) notImplementedMethod(task Entity.Task) []Entity.Result {
	results := []Entity.Result{}
	result, _ := Entity.NewResult(task.Uuid, "not implemented")
	results = append(results, result)
	return results
}
