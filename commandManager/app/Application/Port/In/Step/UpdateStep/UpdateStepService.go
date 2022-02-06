package UpdateStep

import (
	"errors"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
)

type Service struct {
	FindStepPort   TaskOutPort.FindStep
	FindTaskPort   TaskOutPort.FindTask
	UpdateStepPort TaskOutPort.UpdateStepPort
}

func (service Service) Update(command Command) error {
	step, err := service.FindStepPort.Find(command.Uuid)
	if err != nil {
		return errors.New("command not found")
	}
	if command.TaskUuid.Change {
		task, err := service.FindTaskPort.Find(command.TaskUuid.Value)
		if err != nil {
			return errors.New("task not found")
		}
		step.TaskUuid = task.Uuid
		if err != nil {
			return errors.New("internal")
		}
	}
	if command.Sentence.Change {
		step.Sentence = command.Sentence.Value
	}
	err = service.UpdateStepPort.Update(step)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
