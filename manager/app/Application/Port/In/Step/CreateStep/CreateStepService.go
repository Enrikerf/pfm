package CreateStep

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

type Service struct {
	FindTaskPort TaskPort.Find
	SaveStepPort StepPort.Save
}

func (service Service) Create(command Command) (Entity.Step, error) {
	task, err := service.FindTaskPort.Find(command.TaskUuid)
	if err != nil {
		return Entity.Step{}, err
	}
	stepVo, err := ValueObject.NewStepVo(command.Sentence)
	if err != nil {
		return Entity.Step{}, err
	}
	result := Entity.NewStep(task.Uuid, stepVo)
	err = service.SaveStepPort.Save(result)
	if err != nil {
		return Entity.Step{}, err
	}
	return result, nil
}
