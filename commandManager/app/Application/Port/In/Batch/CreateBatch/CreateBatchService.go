package CreateBatch

import (
	ResultOutPorts "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindBatchPort ResultOutPorts.FindBatch
	SaveBatchPort ResultOutPorts.SaveBatchPort
}

func (service Service) Create(command Command) (ResultDomain.Batch, error) {
	task, err := service.FindBatchPort.Find(command.TaskUuid)
	if err != nil {
		return ResultDomain.Batch{}, err
	}
	result := ResultDomain.NewBatch(task.Uuid, []ResultDomain.Result{})
	if err != nil {
		return ResultDomain.Batch{}, err
	}
	err = service.SaveBatchPort.Save(result)
	if err != nil {
		return ResultDomain.Batch{}, err
	}
	return result, nil
}
