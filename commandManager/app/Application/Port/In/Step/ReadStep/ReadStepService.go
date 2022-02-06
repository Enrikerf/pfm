package ReadStep

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindStepPort StepPort.Find
}

func (service Service) Read(query Query) (TaskDomain.Step, error) {
	step, err := service.FindStepPort.Find(query.Uuid)
	if err != nil {
		return TaskDomain.Step{}, errors.New("not found")
	}
	return step, nil
}
