package CreateStep

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindTaskPort TaskPort.Find
	SaveStepPort StepPort.Save
}

func (service Service) Create(command Command) (TaskDomain.Step, error) {
	task, err := service.FindTaskPort.Find(command.TaskUuid)
	if err != nil {
		return TaskDomain.Step{}, err
	}
	result, err := TaskDomain.NewCommand(task.Uuid, command.Name)
	if err != nil {
		return TaskDomain.Step{}, err
	}
	err = service.SaveStepPort.Save(result)
	if err != nil {
		return TaskDomain.Step{}, err
	}
	return result, nil
}
