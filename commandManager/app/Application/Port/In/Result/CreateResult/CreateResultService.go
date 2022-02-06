package CreateResult

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindBatchPort  BatchPort.Find
	SaveResultPort ResultPort.Save
}

func (service Service) Create(command Command) (ResultDomain.Result, error) {
	batch, err := service.FindBatchPort.Find(command.BatchUuid)
	if err != nil {
		return ResultDomain.Result{}, err
	}
	result, err := ResultDomain.NewResult(batch.Uuid, command.Content)
	if err != nil {
		return ResultDomain.Result{}, err
	}
	err = service.SaveResultPort.Save(result)
	if err != nil {
		return ResultDomain.Result{}, err
	}
	return result, nil
}
