package Controller

import (
	"context"
	"fmt"
	taskProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ReadTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/UpdateTask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskController struct {
	SaveTaskUseCase   CreateTask.UseCase
	ListTasksUseCase  ListTasks.UseCase
	ReadTaskUseCase   ReadTask.UseCase
	DeleteTaskUseCase DeleteTask.UseCase
	UpdateTaskUseCase UpdateTask.UseCase
	taskProto.UnimplementedTaskServiceServer
}

func (controller TaskController) CreateTask(ctx context.Context, request *taskProto.CreateTaskRequest) (*taskProto.CreateTaskResponse, error) {
	protoTask := request.GetTaskParams()
	var command CreateTask.Command
	command.Host = protoTask.GetHost()
	command.Port = protoTask.GetPort()
	command.CommandSentences = protoTask.GetCommands()
	command.Mode = protoTask.GetMode().String()
	command.ExecutionMode = protoTask.GetExecutionMode().String()

	task, err := controller.SaveTaskUseCase.Create(command)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	var commandNames []string
	for _, command := range task.Steps {
		commandNames = append(commandNames, command.Sentence)
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

func (controller TaskController) ReadTask(ctx context.Context, request *taskProto.ReadTaskRequest) (*taskProto.ReadTaskResponse, error) {
	var query = ReadTask.Query{Uuid: request.GetTaskUuid()}
	task, err := controller.ReadTaskUseCase.Read(query)
	if err != nil {
		return &taskProto.ReadTaskResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &taskProto.ReadTaskResponse{Task: &taskProto.Task{
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
	cmd := UpdateTask.Command{}
	params := request.GetParams()
	cmd.Uuid = request.GetTaskUuid()
	if params.GetHost() != nil {
		cmd.Host.Change = true
	}
	if params.GetPort() != nil {
		cmd.Port.Change = true
	}
	if params.GetMode() != 0 {
		cmd.Mode.Change = true
	}
	if params.GetStatus() != 0 {
		cmd.Status.Change = true
	}
	if params.GetExecutionMode() != 0 {
		cmd.ExecutionMode.Change = true
	}
	cmd.Host.Value = params.GetHost().GetValue()
	cmd.Port.Value = params.GetPort().GetValue()
	cmd.Mode.Value = params.GetMode().String()
	cmd.Status.Value = params.GetStatus().String()
	cmd.ExecutionMode.Value = params.GetExecutionMode().String()
	err := controller.UpdateTaskUseCase.Update(cmd)
	return &taskProto.UpdateTaskResponse{}, err
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
