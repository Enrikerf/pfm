package CreateResult

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindBatchPort  BatchPort.Find
	SaveResultPort ResultPort.Save
}

func (service Service) Create(command Command) (Entity.Result, error) {
	batch, err := service.FindBatchPort.Find(command.BatchUuid)
	if err != nil {
		return Entity.Result{}, err
	}
	result, err := Entity.NewResult(batch.Uuid, command.Content)
	if err != nil {
		return Entity.Result{}, err
	}
	err = service.SaveResultPort.Save(result)
	if err != nil {
		return Entity.Result{}, err
	}
	return result, nil
}
