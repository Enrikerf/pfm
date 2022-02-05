package ReadCommand

import (
	"errors"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindCommandPort TaskOutPort.FindCommand
}

func (service Service) Read(query Query) (TaskDomain.Command, error) {
	command, err := service.FindCommandPort.Find(query.Uuid)
	if err != nil {
		return TaskDomain.Command{}, errors.New("not found")
	}
	return command, nil
}
