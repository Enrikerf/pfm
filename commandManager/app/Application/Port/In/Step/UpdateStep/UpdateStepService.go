package UpdateStep

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
)

type Service struct {
	FindStepPort   StepPort.Find
	FindTaskPort   TaskPort.Find
	UpdateStepPort StepPort.Update
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
