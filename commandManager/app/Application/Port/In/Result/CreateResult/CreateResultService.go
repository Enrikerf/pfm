package CreateResult

import (
	ResultOutPorts "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindBatchPort ResultOutPorts.FindBatch
	SavePort      ResultOutPorts.SavePort
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
	err = service.SavePort.Save(result)
	if err != nil {
		return ResultDomain.Result{}, err
	}
	return result, nil
}
