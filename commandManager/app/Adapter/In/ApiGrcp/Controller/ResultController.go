package Controller

import (
	"context"
	resultProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/result"
)

type ResultController struct {
	resultProto.UnimplementedResultServiceServer
}

func (controller TaskController) CreateResult(ctx context.Context, request *resultProto.CreateResultRequest) (*resultProto.CreateResultResponse, error) {
	panic("implement me")
}

func (controller TaskController) ReadResult(ctx context.Context, request *resultProto.ReadResultRequest) (*resultProto.ReadResultResponse, error) {
	panic("implement me")
}

func (controller ResultController) UpdateResult(ctx context.Context, request *resultProto.UpdateResultRequest) (*resultProto.UpdateResultResponse, error) {
	panic("implement me")
}

func (controller ResultController) DeleteResult(ctx context.Context, request *resultProto.DeleteResultRequest) (*resultProto.DeleteResultResponse, error) {
	panic("implement me")
}

func (controller ResultController) ListResult(request *resultProto.ListResultRequest, server2 resultProto.ResultService_ListResultServer) error {
	panic("implement me")
}
