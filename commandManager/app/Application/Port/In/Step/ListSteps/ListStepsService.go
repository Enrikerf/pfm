package ListSteps

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindStepByPort StepPort.FindBy
}

func (service Service) List(query Query) []Entity.Step {
	return service.FindStepByPort.FindBy(query)
}
