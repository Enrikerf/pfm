package DeleteResult

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
)

type Service struct {
	DeleteResultPort ResultPort.Delete
}

func (service Service) Delete(command Command) error {
	err := service.DeleteResultPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
