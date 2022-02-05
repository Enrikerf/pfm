package ListCommands

import (
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindCommandByPort TaskOutPort.FindCommandByPort
}

func (service Service) List(query Query) []TaskDomain.Command {
	return service.FindCommandByPort.FindBy(query)
}
