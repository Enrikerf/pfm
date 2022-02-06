package CreateBatch

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Service struct {
	FindBatchPort BatchPort.Find
	SaveBatchPort BatchPort.Save
}

func (service Service) Create(command Command) (Entity.Batch, error) {
	task, err := service.FindBatchPort.Find(command.TaskUuid)
	if err != nil {
		return Entity.Batch{}, err
	}
	result := Entity.NewBatch(task.Uuid, []Entity.Result{})
	if err != nil {
		return Entity.Batch{}, err
	}
	err = service.SaveBatchPort.Save(result)
	if err != nil {
		return Entity.Batch{}, err
	}
	return result, nil
}
