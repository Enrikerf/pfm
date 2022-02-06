package ReadResult

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindResultPort ResultPort.Find
}

func (service Service) Read(query Query) (Result.Result, error) {
	task, err := service.FindResultPort.Find(query.Uuid)
	if err != nil {
		return Result.Result{}, errors.New("not found")
	}
	return task, nil
}
