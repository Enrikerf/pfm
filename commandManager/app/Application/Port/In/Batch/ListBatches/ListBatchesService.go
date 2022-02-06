package ListBatches

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindBatchesByPort BatchPort.FindBy
}

func (service Service) List(query Query) []Result.Batch {
	return service.FindBatchesByPort.FindBy(query)
}
