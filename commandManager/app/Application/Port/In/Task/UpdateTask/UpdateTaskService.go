package UpdateTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindTaskPort   TaskPort.Find
	UpdateTaskPort TaskPort.Update
}

func (service Service) Update(command Command) error {
	//TODO: error types from ports and error types to return from domain
	task, err := service.FindTaskPort.Find(command.Uuid)
	if err != nil {
		return errors.New("not found")
	}
	if command.Host.Change {
		task.Host = command.Host.Value
	}
	if command.Port.Change {
		task.Port = command.Port.Value
	}
	if command.Mode.Change {
		task.Mode, err = Task.GetTaskMode(command.Mode.Value)
	}
	if command.Status.Change {
		task.Status, err = Task.GetStatus(command.Status.Value)
	}
	if command.ExecutionMode.Change {
		task.ExecutionMode, err = Task.GetExecutionMode(command.ExecutionMode.Value)
	}

	err = service.UpdateTaskPort.Update(task)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
