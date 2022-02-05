package DeleteCommand

import (
	"errors"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
)

type Service struct {
	DeleteCommandPort TaskOutPort.DeleteCommand
}

func (service Service) Delete(command Command) error {
	err := service.DeleteCommandPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
