package Communication

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"google.golang.org/grpc"
	"io"
)

type Adapter struct {
}

func (adapter Adapter) Communicate(task Task.Task) Result.Batch {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.GetHost().GetValue()+":"+task.GetPort().GetValue(), options)
	var results []Content.Content
	if err != nil {
		fmt.Printf(err.Error())
		result, _ := Content.NewContent(err.Error())
		results = append(results, result)
	} else {
		defer func(connection *grpc.ClientConn) {
			err := connection.Close()
			if err != nil {

			}
		}(connection)
		client := call.NewCallServiceClient(connection)
		switch task.GetCommunicationMode() {
		case CommunicationMode.Unary:
			results = adapter.doUnaryCall(task, client)
		case CommunicationMode.ServerStream:
			results = adapter.doServerStream(task, client)
		case CommunicationMode.ClientStream:
			results = adapter.doClientStream(task, client)
		case CommunicationMode.Bidirectional:
			results = adapter.doBidirectional(task, client)
		}
	}
	batch := Result.NewBatch(task.GetId())
	batch.SetResultsFromContent(results)
	return batch
}

func (adapter Adapter) doServerStream(task Task.Task, client call.CallServiceClient) []Content.Content {
	var results []Content.Content
	request := &call.CallRequest{
		Step: task.GetSteps()[0].GetSentence(),
	}
	responseStream, err := client.CallServerStream(context.Background(), request)

	if err != nil {
		result, _ := Content.NewContent(err.Error())
		results = append(results, result)
	} else {
		for {
			msg, err := responseStream.Recv()
			if err == io.EOF {
				result, _ := Content.NewContent(err.Error())
				results = append(results, result)
				break
			}
			if err != nil {
				result, _ := Content.NewContent(err.Error())
				results = append(results, result)
				break
			}
			result, _ := Content.NewContent(msg.GetResult())
			results = append(results, result)
		}
	}
	return results
}

func (adapter Adapter) doBidirectional(task Task.Task, client call.CallServiceClient) []Content.Content {
	var results []Content.Content
	fmt.Println("find maximum")
	stream, err := client.CallBidirectional(context.Background())
	if err != nil {
		result, _ := Content.NewContent(err.Error())
		results = append(results, result)
	} else {
		// we send a bunch of messages to de client go routine
		resultsChannel1 := make(chan string)
		go sendInStream(task, stream, resultsChannel1)
		// we receive a bunch of messages form the client go routine
		resultsChannel2 := make(chan string)
		go receiveServerStream(stream, resultsChannel2)
		//TODO: missing wait channels
		for i := range resultsChannel1 {
			result, _ := Content.NewContent(i)
			results = append(results, result)
		}
		for i := range resultsChannel2 {
			result, _ := Content.NewContent(i)
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
		fmt.Printf("received: %v\n", response)
	}
	close(resultsChannel)
}

func sendInStream(task Task.Task, stream call.CallService_CallBidirectionalClient, resultsChannel chan string) {
	for _, step := range task.GetSteps() {
		fmt.Printf("sending message %v\n", step)
		callRequest := call.CallRequest{
			Step: step.GetSentence(),
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

func (adapter Adapter) doClientStream(task Task.Task, client call.CallServiceClient) []Content.Content {
	var results []Content.Content
	stream, err := client.CallClientStream(context.Background())
	if err != nil {
		result, _ := Content.NewContent(err.Error())
		results = append(results, result)
	} else {
		for _, step := range task.GetSteps() {
			fmt.Printf("sending: %v\n", step.GetSentence())
			err := stream.Send(&call.CallRequest{
				Step: step.GetSentence(),
			})
			if err != nil {
				result, _ := Content.NewContent(err.Error())
				results = append(results, result)
			}
		}
		response, err := stream.CloseAndRecv()
		if err != nil {
			result, _ := Content.NewContent(err.Error())
			results = append(results, result)
		}
		result, _ := Content.NewContent(response.GetResult())
		results = append(results, result)
	}
	return results
}

func (adapter Adapter) doUnaryCall(task Task.Task, client call.CallServiceClient) []Content.Content {
	var results []Content.Content
	var result Content.Content
	callRequest := call.CallRequest{
		Step: task.GetSteps()[0].GetSentence(),
	}
	callResponse, err := client.CallUnary(context.Background(), &callRequest)
	if err != nil {
		result, _ = Content.NewContent(err.Error())
	} else {
		result, _ = Content.NewContent(callResponse.GetResult())
	}
	return append(results, result)
}
