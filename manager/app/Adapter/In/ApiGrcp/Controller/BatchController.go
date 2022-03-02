package Controller

import (
	"context"
	"fmt"
	batchProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/batch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/CreateBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/DeleteBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ListBatches"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ReadBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/UpdateBatch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BatchController struct {
	CreateBatchUseCase CreateBatch.UseCase
	ReadBatchUseCase   ReadBatch.UseCase
	UpdateBatchUseCase UpdateBatch.UseCase
	DeleteBatchUseCase DeleteBatch.UseCase
	ListBatchesUseCase ListBatches.UseCase
	batchProto.UnimplementedBatchServiceServer
}

func (controller BatchController) CreateBatch(ctx context.Context, request *batchProto.CreateBatchRequest) (*batchProto.CreateBatchResponse, error) {
	protoBatch := request.GetBatchParams()
	var command CreateBatch.Command
	command.TaskUuid = protoBatch.GetTaskUuid()
	batch, err := controller.CreateBatchUseCase.Create(command)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	newBatch := batchProto.Batch{
		Uuid:      batch.Uuid.String(),
		TaskUuid:  batch.TaskUuid.String(),
		CreatedAt: batch.CreatedAt.String(),
		UpdatedAt: batch.UpdatedAt.String(),
	}
	return &batchProto.CreateBatchResponse{Batch: &newBatch}, nil
}

func (controller BatchController) ReadBatch(ctx context.Context, request *batchProto.ReadBatchRequest) (*batchProto.ReadBatchResponse, error) {
	var query = ReadBatch.Query{Uuid: request.GetBatchUuid()}
	batch, err := controller.ReadBatchUseCase.Read(query)
	if err != nil {
		return &batchProto.ReadBatchResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}
	return &batchProto.ReadBatchResponse{Batch: &batchProto.Batch{
		Uuid:     batch.Uuid.String(),
		TaskUuid: batch.TaskUuid.String(),
	}}, nil
}

func (controller BatchController) UpdateBatch(ctx context.Context, request *batchProto.UpdateBatchRequest) (*batchProto.UpdateBatchResponse, error) {
	cmd := UpdateBatch.Command{}
	params := request.GetParams()
	cmd.Uuid = request.GetBatchUuid()
	cmd.TaskUuid.Change = true
	cmd.TaskUuid.Value = params.GetTaskUuid()
	err := controller.UpdateBatchUseCase.Update(cmd)
	return &batchProto.UpdateBatchResponse{}, err
}

func (controller BatchController) DeleteBatch(ctx context.Context, request *batchProto.DeleteBatchRequest) (*batchProto.DeleteBatchResponse, error) {
	var command = DeleteBatch.Command{Uuid: request.GetBatchUuid()}
	err := controller.DeleteBatchUseCase.Delete(command)
	if err != nil {
		return &batchProto.DeleteBatchResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &batchProto.DeleteBatchResponse{}, nil
}

func (controller BatchController) ListBatches(ctx context.Context, request *batchProto.ListBatchesRequest) (*batchProto.ListBatchesResponse, error) {
	query := ListBatches.Query{}
	if request.GetFilters().GetTaskUuid() != nil {
		query.TaskUuid.Change = true
	}
	query.TaskUuid.Value = request.GetFilters().GetTaskUuid().Value

	batches := controller.ListBatchesUseCase.List(query)
	if batches == nil {
		return &batchProto.ListBatchesResponse{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error"),
		)
	}
	batchProtoArray := []*batchProto.Batch{}
	for _, batch := range batches {
		batchProtoArray = append(batchProtoArray, &batchProto.Batch{
			Uuid:      batch.Uuid.String(),
			TaskUuid:  batch.TaskUuid.String(),
			CreatedAt: batch.CreatedAt.String(),
			UpdatedAt: batch.UpdatedAt.String(),
		})
	}
	return &batchProto.ListBatchesResponse{Batches: batchProtoArray}, nil
}
