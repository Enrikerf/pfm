package UpdateCommand

import (
	"errors"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
)

type Service struct {
	FindCommandPort   TaskOutPort.FindCommand
	FindTaskPort      TaskOutPort.Find
	UpdateCommandPort TaskOutPort.UpdateCommandPort
}

func (service Service) Update(command Command) error {
	commandDomain, err := service.FindCommandPort.Find(command.Uuid)
	if err != nil {
		return errors.New("command not found")
	}
	if command.TaskUuid.Change {
		task, err := service.FindTaskPort.Find(command.TaskUuid.Value)
		if err != nil {
			return errors.New("task not found")
		}
		commandDomain.TaskUuid = task.Uuid
		if err != nil {
			return errors.New("internal")
		}
	}
	if command.Name.Change {
		commandDomain.Name = command.Name.Value
	}
	err = service.UpdateCommandPort.Update(commandDomain)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
