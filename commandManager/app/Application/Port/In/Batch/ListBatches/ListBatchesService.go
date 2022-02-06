package ListBatches

import (
	ResultOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindBatchesByPort ResultOutPort.FindBatchesByPort
}

func (service Service) List(query Query) []Result.Batch {
	return service.FindBatchesByPort.FindBy(query)
}
