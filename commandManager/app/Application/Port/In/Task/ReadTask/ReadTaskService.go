package ReadTask

import (
	"errors"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindTaskPort TaskOutPort.FindTask
}

func (service Service) Read(query Query) (Task.Task, error) {
	task, err := service.FindTaskPort.Find(query.Uuid)
	if err != nil {
		return Task.Task{}, errors.New("not found")
	}
	return task, nil
}
