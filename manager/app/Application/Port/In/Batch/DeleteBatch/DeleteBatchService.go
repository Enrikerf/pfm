package DeleteBatch

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
)

type Service struct {
	DeleteBatchPort BatchPort.Delete
}

func (service Service) Delete(command Command) error {
	err := service.DeleteBatchPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
