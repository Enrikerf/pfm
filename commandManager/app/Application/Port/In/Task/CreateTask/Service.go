package CreateTask

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model"

type Service struct{}

func (service Service) Create(command Command) (Model.Task, error) {
	var task, _ = Model.NewTask(command.Host, command.Port, command.Command, command.Mode)
	return task, nil
}
