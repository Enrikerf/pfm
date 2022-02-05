package UpdateResult

import (
	"errors"
	ResultOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	"github.com/google/uuid"
)

type Service struct {
	FindPort      ResultOutPort.Find
	FindBatchPort ResultOutPort.FindBatch
	UpdatePort    ResultOutPort.UpdatePort
}

func (service Service) Update(command Command) error {
	//TODO: error types from ports and error types to return from domain
	result, err := service.FindPort.Find(command.Uuid)
	if err != nil {
		return errors.New("not found")
	}
	batch, err := service.FindBatchPort.Find(command.BatchUuid.Value)
	if err != nil {
		return errors.New("not found")
	}
	if command.Content.Change {
		result.Content = command.Content.Value
	}
	if command.BatchUuid.Change {
		result.BatchUuid, err = uuid.Parse(batch.Uuid.String())
		if err != nil {
			return errors.New("internal")
		}
	}
	err = service.UpdatePort.Update(result)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
