package ListBatches

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindBatchesByPort BatchPort.FindBy
}

func (service Service) List(query Query) []Entity.Batch {
	return service.FindBatchesByPort.FindBy(query)
}
