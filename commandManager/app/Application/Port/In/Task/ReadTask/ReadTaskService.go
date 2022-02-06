package ReadTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindTaskPort TaskPort.Find
}

func (service Service) Read(query Query) (Task.Task, error) {
	task, err := service.FindTaskPort.Find(query.Uuid)
	if err != nil {
		return Task.Task{}, errors.New("not found")
	}
	return task, nil
}
