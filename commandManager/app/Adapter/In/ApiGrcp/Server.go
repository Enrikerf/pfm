package ApiGrcp

import (
	"context"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/proto"
)

type Server struct {
	proto.UnimplementedCommandServiceServer
}

func (server Server) CreateCommand(ctx context.Context, request *proto.CreateCommandRequest) (*proto.CreateCommandResponse, error) {
	panic("implement me")
}

func (server Server) ReadCommand(ctx context.Context, request *proto.ReadCommandRequest) (*proto.ReadCommandResponse, error) {
	panic("implement me")
}

func (server Server) UpdateCommand(ctx context.Context, request *proto.UpdateCommandRequest) (*proto.UpdateCommandResponse, error) {
	panic("implement me")
}

func (server Server) DeleteCommand(ctx context.Context, request *proto.DeleteCommandRequest) (*proto.DeleteCommandResponse, error) {
	panic("implement me")
}

func (server Server) ListCommand(request *proto.ListCommandRequest, server2 proto.CommandService_ListCommandServer) error {
	panic("implement me")
}
