package Controller

import (
	"context"
	"fmt"
	commandProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/command"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Command/CreateCommand"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Command/DeleteCommand"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Command/ListCommands"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Command/ReadCommand"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Command/UpdateCommand"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommandController struct {
	CreateCommandUseCase CreateCommand.UseCase
	ReadCommandUseCase   ReadCommand.UseCase
	UpdateCommandUseCase UpdateCommand.UseCase
	DeleteCommandUseCase DeleteCommand.UseCase
	ListCommandsUseCase  ListCommands.UseCase
	commandProto.UnimplementedCommandServiceServer
}

func (controller CommandController) CreateCommand(ctx context.Context, request *commandProto.CreateCommandRequest) (*commandProto.CreateCommandResponse, error) {
	protoCommand := request.GetCommandParams()
	var command CreateCommand.Command
	command.TaskUuid = protoCommand.GetTaskUuid()
	command.Name = protoCommand.GetName()
	commandDomain, err := controller.CreateCommandUseCase.Create(command)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	newCommand := commandProto.Command{
		Uuid:     commandDomain.Uuid.String(),
		TaskUuid: commandDomain.TaskUuid.String(),
		Name:     commandDomain.Name,
	}
	return &commandProto.CreateCommandResponse{Command: &newCommand}, nil
}

func (controller CommandController) ReadCommand(ctx context.Context, request *commandProto.ReadCommandRequest) (*commandProto.ReadCommandResponse, error) {
	var query = ReadCommand.Query{Uuid: request.GetCommandUuid()}
	command, err := controller.ReadCommandUseCase.Read(query)
	if err != nil {
		return &commandProto.ReadCommandResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}
	return &commandProto.ReadCommandResponse{Command: &commandProto.Command{
		Uuid:     command.Uuid.String(),
		TaskUuid: command.TaskUuid.String(),
		Name:     command.Name,
	}}, nil
}

func (controller CommandController) UpdateCommand(ctx context.Context, request *commandProto.UpdateCommandRequest) (*commandProto.UpdateCommandResponse, error) {
	cmd := UpdateCommand.Command{}
	params := request.GetCommandParams()
	cmd.Uuid = request.GetCommandUuid()
	if params.GetTaskUuid() != nil {
		cmd.TaskUuid.Change = true
	}
	if params.GetName() != nil {
		cmd.Name.Change = true
	}
	cmd.TaskUuid.Value = params.GetTaskUuid().Value
	cmd.Name.Value = params.GetName().Value
	err := controller.UpdateCommandUseCase.Update(cmd)
	return &commandProto.UpdateCommandResponse{}, err
}

func (controller CommandController) DeleteCommand(ctx context.Context, request *commandProto.DeleteCommandRequest) (*commandProto.DeleteCommandResponse, error) {
	var command = DeleteCommand.Command{Uuid: request.GetCommandUuid()}
	err := controller.DeleteCommandUseCase.Delete(command)
	if err != nil {
		return &commandProto.DeleteCommandResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &commandProto.DeleteCommandResponse{}, nil
}

func (controller CommandController) ListCommands(ctx context.Context, request *commandProto.ListCommandsRequest) (*commandProto.ListCommandsResponse, error) {
	commands := controller.ListCommandsUseCase.List(ListCommands.Query{})
	if commands == nil {
		return &commandProto.ListCommandsResponse{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error"),
		)
	}
	commandProtoArray := []*commandProto.Command{}
	for _, command := range commands {
		commandProtoArray = append(commandProtoArray, &commandProto.Command{
			Uuid:     command.Uuid.String(),
			TaskUuid: command.TaskUuid.String(),
			Name:     command.Name,
		})
	}
	return &commandProto.ListCommandsResponse{Commands: commandProtoArray}, nil
}
