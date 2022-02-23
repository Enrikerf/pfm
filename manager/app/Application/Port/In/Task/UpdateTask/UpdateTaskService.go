package UpdateTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Call/Manual"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

type Service struct {
	FindTaskPort   TaskPort.Find
	UpdateTaskPort TaskPort.Update
	ManualCallPort Manual.UseCase
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
		task.Mode, err = ValueObject.GetTaskMode(command.Mode.Value)
	}
	if command.Status.Change {
		task.Status, err = ValueObject.GetStatus(command.Status.Value)
	}
	if command.ExecutionMode.Change {
		task.ExecutionMode, err = ValueObject.GetExecutionMode(command.ExecutionMode.Value)
	}

	err = service.UpdateTaskPort.Update(task)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
