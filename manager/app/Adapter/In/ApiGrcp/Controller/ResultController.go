package Controller

import (
	"context"
	"fmt"
	resultProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/result"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/CreateBatchAndFill"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/GetBatchResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/StreamResults"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"time"
)

type ResultController struct {
	ExecuteTaskManuallyUseCase CreateBatchAndFill.UseCase
	GetBatchResultsUseCase     GetBatchResults.UseCase
	StreamResultsUseCase       StreamResults.UseCase
	resultProto.UnimplementedResultServiceServer
}

func (controller ResultController) CreateBatchAndFill(
	ctx context.Context,
	request *resultProto.CreateBatchAndFillRequest,
) (*resultProto.CreateBatchAndFillResponse, error) {
	var command CreateBatchAndFill.Command
	command.TaskUuid = request.GetTaskUuid()
	batch, err := controller.ExecuteTaskManuallyUseCase.Execute(command)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &resultProto.CreateBatchAndFillResponse{
		BatchUuid: batch.GetId().GetUuidString(),
	}, nil
}

func (controller ResultController) GetBatchResults(
	ctx context.Context,
	request *resultProto.GetBatchResultsRequest,
) (*resultProto.ListResultsResponse, error) {
	batchId, err := Result.LoadBatchIdFromString(request.GetBatchUuid())
	if err != nil {
		return nil, err
	}
	results := controller.GetBatchResultsUseCase.List(GetBatchResults.Query{BatchId: batchId})
	var resultProtoArray []*resultProto.Result
	for _, result := range results {
		resultProtoArray = append(resultProtoArray, &resultProto.Result{
			Uuid:      result.GetId().GetUuidString(),
			BatchUuid: result.GetBatchId().GetUuidString(),
			Content:   result.GetContent().GetValue(),
			CreatedAt: result.GetCreateAt().Format(time.RFC3339),
		})
	}
	return &resultProto.ListResultsResponse{Results: resultProtoArray}, nil
}

func (controller ResultController) GetTaskBatches(
	ctx context.Context,
	request *resultProto.GetTaskBatchesRequest,
) (*resultProto.ListBatchesResponse, error) {
	return &resultProto.ListBatchesResponse{}, nil
}

func (controller ResultController) StreamResults(
	request *resultProto.StreamResultsRequest,
	stream resultProto.ResultService_StreamResultsServer,
) error {
	fmt.Printf("streaming results %v\n", request)
	var lastId Result.Id
	batchUuid, err := Result.LoadBatchIdFromString(request.GetBatchUuid())
	if err != nil {
		return err
	}
	for {
		results, err := controller.StreamResultsUseCase.Stream(StreamResults.Query{
			BatchUuid: batchUuid,
			LastId:    lastId,
		})
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}
		if results == nil {
			fmt.Printf("Task Done")
			return nil
		}
		if len(results) > 0 {
			var resultProtoArray []*resultProto.Result
			for _, result := range results {
				resultProtoArray = append(resultProtoArray, &resultProto.Result{
					Uuid:      result.GetId().GetUuidString(),
					BatchUuid: result.GetBatchId().GetUuidString(),
					Content:   result.GetContent().GetValue(),
					CreatedAt: result.GetCreateAt().Format(time.RFC3339),
				})
			}
			err := stream.Send(&resultProto.StreamResultsResponse{Results: resultProtoArray})
			if err != nil {
				return err
			}
			lastId = results[len(results)-1].GetId()
		}
		time.Sleep(100 * time.Millisecond)
	}
}
