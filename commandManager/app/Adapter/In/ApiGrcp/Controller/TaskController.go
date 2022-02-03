package Controller

import (
	"context"
	"fmt"
	taskProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ShowTask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskController struct {
	SaveTaskUseCase   CreateTask.UseCase
	ListTasksUseCase  ListTasks.UseCase
	ShowTaskUseCase   ShowTask.UseCase
	DeleteTaskUseCase DeleteTask.UseCase
	taskProto.UnimplementedTaskServiceServer
}

func (controller TaskController) CreateTask(ctx context.Context, request *taskProto.CreateTaskRequest) (*taskProto.CreateTaskResponse, error) {
	protoTask := request.GetTask()
	var command CreateTask.Command
	command.Host = protoTask.Host
	command.Port = protoTask.Port
	command.Commands = protoTask.GetCommands()
	command.Mode = protoTask.Mode
	command.ExecutionMode = protoTask.ExecutionMode

	task, err := controller.SaveTaskUseCase.Create(command)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	var commandNames []string
	for _, command := range task.Commands {
		commandNames = append(commandNames, command.Name)
	}
	newTask := taskProto.Task{
		Uuid:          task.Uuid.String(),
		Host:          task.Host,
		Port:          task.Port,
		Commands:      commandNames,
		Mode:          task.Mode.String(),
		Status:        task.Status.String(),
		ExecutionMode: task.ExecutionMode.String(),
	}
	return &taskProto.CreateTaskResponse{Task: &newTask}, nil
}

func (controller TaskController) ShowTask(ctx context.Context, request *taskProto.ShowTaskRequest) (*taskProto.ShowTaskResponse, error) {

	var query = ShowTask.Query{Uuid: request.GetTaskUuid()}
	task, err := controller.ShowTaskUseCase.Show(query)
	if err != nil {
		return &taskProto.ShowTaskResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &taskProto.ShowTaskResponse{Task: &taskProto.Task{
		Uuid:          task.Uuid.String(),
		Host:          task.Host,
		Port:          task.Port,
		Commands:      nil,
		Mode:          task.Mode.String(),
		Status:        task.Status.String(),
		ExecutionMode: task.ExecutionMode.String(),
	}}, nil
}

func (controller TaskController) UpdateTask(ctx context.Context, request *taskProto.UpdateTaskRequest) (*taskProto.UpdateTaskResponse, error) {
	panic("implement me")
}

func (controller TaskController) DeleteTask(ctx context.Context, request *taskProto.DeleteTaskRequest) (*taskProto.DeleteTaskResponse, error) {
	var command = DeleteTask.Command{Uuid: request.GetTaskUuid()}
	err := controller.DeleteTaskUseCase.Delete(command)
	if err != nil {
		return &taskProto.DeleteTaskResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &taskProto.DeleteTaskResponse{}, nil
}

func (controller TaskController) ListTasks(ctx context.Context, in *taskProto.ListTasksRequest) (*taskProto.ListTasksResponse, error) {
	tasks := controller.ListTasksUseCase.List(ListTasks.Query{})
	if tasks == nil {
		return &taskProto.ListTasksResponse{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error"),
		)
	}
	tasksProtoArray := []*taskProto.Task{}
	for _, task := range tasks {
		tasksProtoArray = append(tasksProtoArray, &taskProto.Task{
			Uuid:          task.Uuid.String(),
			Host:          task.Host,
			Port:          task.Port,
			Commands:      nil,
			Mode:          task.Mode.String(),
			Status:        task.Status.String(),
			ExecutionMode: task.ExecutionMode.String(),
		})
	}
	return &taskProto.ListTasksResponse{Tasks: tasksProtoArray}, nil
}
