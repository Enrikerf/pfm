package DeleteResult

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
)

type Service struct {
	DeleteTaskPort Result.DeleteResult
}

func (service Service) Delete(command Command) error {
	err := service.DeleteTaskPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
