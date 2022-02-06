package ListTasks

import (
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindTasksByPort TaskOutPort.FindTasksByPort
}

func (service Service) List(query Query) []Task.Task {
	return service.FindTasksByPort.FindBy(query)
}
