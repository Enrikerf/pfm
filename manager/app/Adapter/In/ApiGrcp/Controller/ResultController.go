package Controller

import (
	"context"
	"fmt"
	resultProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/result"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/CreateResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/DeleteResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ListResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ReadResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/StreamResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/UpdateResult"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ResultController struct {
	CreateResultUseCase  CreateResult.UseCase
	ReadResultUseCase    ReadResult.UseCase
	UpdateResultUseCase  UpdateResult.UseCase
	DeleteResultUseCase  DeleteResult.UseCase
	ListResultsUseCase   ListResults.UseCase
	StreamResultsUseCase StreamResults.UseCase
	resultProto.UnimplementedResultServiceServer
}

func (controller ResultController) CreateResult(ctx context.Context, request *resultProto.CreateResultRequest) (*resultProto.CreateResultResponse, error) {
	protoResult := request.GetResultParams()
	var command CreateResult.Command
	command.Content = protoResult.GetContent()
	command.BatchUuid = protoResult.GetBatchUuid()
	result, err := controller.CreateResultUseCase.Create(command)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	newResult := resultProto.Result{
		Uuid:      result.Uuid.String(),
		Content:   result.Content,
		BatchUuid: result.BatchUuid.String(),
	}
	return &resultProto.CreateResultResponse{Result: &newResult}, nil
}

func (controller ResultController) ReadResult(ctx context.Context, request *resultProto.ReadResultRequest) (*resultProto.ReadResultResponse, error) {
	var query = ReadResult.Query{Uuid: request.GetResultUuid()}
	result, err := controller.ReadResultUseCase.Read(query)
	if err != nil {
		return &resultProto.ReadResultResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}
	return &resultProto.ReadResultResponse{Result: &resultProto.Result{
		Uuid:      result.Uuid.String(),
		Content:   result.Content,
		BatchUuid: result.BatchUuid.String(),
	}}, nil
}

func (controller ResultController) UpdateResult(ctx context.Context, request *resultProto.UpdateResultRequest) (*resultProto.UpdateResultResponse, error) {
	cmd := UpdateResult.Command{}
	params := request.GetResultParams()
	cmd.Uuid = request.GetResultUuid()
	if params.GetContent() != nil {
		cmd.Content.Change = true
	}
	if params.GetBatchUuid() != nil {
		cmd.BatchUuid.Change = true
	}
	cmd.Content.Value = params.GetContent().GetValue()
	cmd.BatchUuid.Value = params.GetBatchUuid().GetValue()
	err := controller.UpdateResultUseCase.Update(cmd)
	return &resultProto.UpdateResultResponse{}, err
}

func (controller ResultController) DeleteResult(ctx context.Context, request *resultProto.DeleteResultRequest) (*resultProto.DeleteResultResponse, error) {
	var command = DeleteResult.Command{Uuid: request.GetResultUuid()}
	err := controller.DeleteResultUseCase.Delete(command)
	if err != nil {
		return &resultProto.DeleteResultResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &resultProto.DeleteResultResponse{}, nil
}

func (controller ResultController) ListResult(ctx context.Context, request *resultProto.ListResultRequest) (*resultProto.ListResultResponse, error) {
	results := controller.ListResultsUseCase.List(ListResults.Query{})
	if results == nil {
		return &resultProto.ListResultResponse{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error"),
		)
	}
	resultProtoArray := []*resultProto.Result{}
	for _, result := range results {
		resultProtoArray = append(resultProtoArray, &resultProto.Result{
			Uuid:      result.Uuid.String(),
			BatchUuid: result.BatchUuid.String(),
			Content:   result.Content,
		})
	}
	return &resultProto.ListResultResponse{Results: resultProtoArray}, nil
}

func (controller ResultController) StreamResults(request *resultProto.StreamResultsRequest, stream resultProto.ResultService_StreamResultsServer) error {
	fmt.Printf("streaming results %v\n", request)
	lastDate, _ := time.Parse(time.RFC3339, "1000-01-01")
	batchUuid, err := uuid.Parse(request.GetBatchUuid())
	if err != nil {
		return nil
	}
	for {
		results, finish := controller.StreamResultsUseCase.Stream(StreamResults.Query{
			BatchUuid: batchUuid,
			LastDate:  lastDate,
		})
		if finish {
			fmt.Printf("finish streaming results \n")
			return nil
		}
		if len(results) > 0 {
			resultProtoArray := []*resultProto.Result{}
			for _, result := range results {
				resultProtoArray = append(resultProtoArray, &resultProto.Result{
					Uuid:      result.Uuid.String(),
					BatchUuid: result.BatchUuid.String(),
					Content:   result.Content,
					CreatedAt: result.CreatedAt.String(),
					UpdatedAt: result.UpdatedAt.String(),
				})
			}
			response := &resultProto.StreamResultsResponse{Results: resultProtoArray}
			err := stream.Send(response)
			if err != nil {
				return nil
			}
			lastDate = results[len(results)-1].CreatedAt
		}
		//time.Sleep(100 * time.Millisecond)
	}
}
