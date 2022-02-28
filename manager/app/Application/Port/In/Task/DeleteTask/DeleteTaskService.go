package DeleteTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
)

type Service struct {
	DeleteTaskPort TaskPort.Delete
}

func (service Service) Delete(command Command) error {
	err := service.DeleteTaskPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
