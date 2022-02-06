package ListSteps

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindStepByPort StepPort.FindBy
}

func (service Service) List(query Query) []TaskDomain.Step {
	return service.FindStepByPort.FindBy(query)
}
