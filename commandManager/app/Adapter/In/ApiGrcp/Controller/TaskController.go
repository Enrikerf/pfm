package Controller

import (
	"context"
	"fmt"
	taskProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
)

type TaskController struct {
	SaveTaskUseCase CreateTask.UseCase
	taskProto.UnimplementedTaskServiceServer
}

func (controller TaskController) CreateTask(ctx context.Context, request *taskProto.CreateTaskRequest) (*taskProto.CreateTaskResponse, error) {
	protoTask := request.GetTask()
	var command CreateTask.Command
	command.Host = protoTask.Host
	command.Port = protoTask.Port
	command.Command = protoTask.Command
	command.Mode = protoTask.Mode

	task, err := controller.SaveTaskUseCase.Create(command)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	newTask := taskProto.Task{
		Uuid:    task.Uuid.String(),
		Host:    task.Host,
		Port:    task.Port,
		Command: task.Command,
		Mode:    task.Mode.String(),
		Status:  task.Status.String(),
	}
	return &taskProto.CreateTaskResponse{Task: &newTask}, nil
}

func (controller TaskController) ReadTask(ctx context.Context, request *taskProto.ReadTaskRequest) (*taskProto.ReadTaskResponse, error) {
	panic("implement me")
}

func (controller TaskController) UpdateTask(ctx context.Context, request *taskProto.UpdateTaskRequest) (*taskProto.UpdateTaskResponse, error) {
	panic("implement me")
}

func (controller TaskController) DeleteTask(ctx context.Context, request *taskProto.DeleteTaskRequest) (*taskProto.DeleteTaskResponse, error) {
	panic("implement me")
}

func (controller TaskController) ListTask(request *taskProto.ListTaskRequest, server2 taskProto.TaskService_ListTaskServer) error {
	panic("implement me")
}
