package ListTasks

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindTasksByPort TaskPort.FindBy
}

func (service Service) List(query Query) []Entity.Task {
	return service.FindTasksByPort.FindBy(query)
}
