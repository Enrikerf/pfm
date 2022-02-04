package DeleteTask

import (
	"errors"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
)

type Service struct {
	DeleteTaskPort TaskOutPort.Delete
}

func (service Service) Delete(command Command) error {
	err := service.DeleteTaskPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
