package UpdateResult

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
)

type Service struct {
	FindResultPort   ResultPort.Find
	FindBatchPort    BatchPort.Find
	UpdateResultPort ResultPort.Update
}

func (service Service) Update(command Command) error {
	//TODO: error types from ports and error types to return from domain
	result, err := service.FindResultPort.Find(command.Uuid)
	if err != nil {
		return errors.New("not found")
	}
	if command.Content.Change {
		result.Content = command.Content.Value
	}
	if command.BatchUuid.Change {
		batch, err := service.FindBatchPort.Find(command.BatchUuid.Value)
		if err != nil {
			return errors.New("not found")
		}
		result.BatchUuid = batch.Uuid
		if err != nil {
			return errors.New("internal")
		}
	}
	err = service.UpdateResultPort.Update(result)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
