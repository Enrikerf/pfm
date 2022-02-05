package ShowResult

import (
	"errors"
	ResultFindPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindByPort ResultFindPort.Find
}

func (service Service) Show(query Query) (Result.Result, error) {
	task, err := service.FindByPort.Find(query.Uuid)
	if err != nil {
		return Result.Result{}, errors.New("not found")
	}
	return task, nil
}
