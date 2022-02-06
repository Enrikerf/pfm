package ReadResult

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindResultPort ResultPort.Find
}

func (service Service) Read(query Query) (Entity.Result, error) {
	task, err := service.FindResultPort.Find(query.Uuid)
	if err != nil {
		return Entity.Result{}, errors.New("not found")
	}
	return task, nil
}
