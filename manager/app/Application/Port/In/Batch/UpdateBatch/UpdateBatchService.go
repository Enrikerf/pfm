package UpdateBatch

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
)

type Service struct {
	FindBatchPort   BatchPort.Find
	FindTaskPort    TaskPort.Find
	UpdateBatchPort BatchPort.Update
}

func (service Service) Update(command Command) error {
	batch, err := service.FindBatchPort.Find(command.Uuid)
	if err != nil {
		return errors.New("batch not found")
	}
	if command.TaskUuid.Change {
		task, err := service.FindTaskPort.Find(command.TaskUuid.Value)
		if err != nil {
			return errors.New("task not found")
		}
		batch.TaskUuid = task.Uuid
		if err != nil {
			return errors.New("internal")
		}
	}
	err = service.UpdateBatchPort.Update(batch)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
