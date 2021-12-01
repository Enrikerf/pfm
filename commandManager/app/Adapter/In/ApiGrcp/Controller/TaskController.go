package Controller

import (
	"fmt"
	taskProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
)

type TaskController struct {
	SaveTaskUseCase CreateTask.UseCase
}

func (taskController TaskController) PostTask(protoTask *taskProto.Task) (*taskProto.Task, error) {
	var command CreateTask.Command
	command.Host = protoTask.Host
	command.Port = protoTask.Port
	command.Command = protoTask.Command
	command.Mode = protoTask.Mode
	command.Status = protoTask.Status

	task, err := taskController.SaveTaskUseCase.Save(command)
	if err != nil {
		return protoTask, fmt.Errorf("error")
	}
	protoTask.Id = task.Id.String()
	return protoTask, nil
}
