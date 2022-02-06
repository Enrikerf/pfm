package ListSteps

import (
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindStepByPort TaskOutPort.FindStepByPort
}

func (service Service) List(query Query) []TaskDomain.Step {
	return service.FindStepByPort.FindBy(query)
}
