package DeleteBatch

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
)

type Service struct {
	DeleteBatchPort Result.DeleteBatch
}

func (service Service) Delete(command Command) error {
	err := service.DeleteBatchPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
