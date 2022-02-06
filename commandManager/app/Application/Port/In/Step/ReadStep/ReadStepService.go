package ReadStep

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindStepPort StepPort.Find
}

func (service Service) Read(query Query) (Entity.Step, error) {
	step, err := service.FindStepPort.Find(query.Uuid)
	if err != nil {
		return Entity.Step{}, errors.New("not found")
	}
	return step, nil
}
