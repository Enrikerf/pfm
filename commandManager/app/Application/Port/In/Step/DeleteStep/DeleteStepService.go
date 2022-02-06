package DeleteStep

import (
	"errors"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
)

type Service struct {
	DeleteStepPort TaskOutPort.DeleteStep
}

func (service Service) Delete(command Command) error {
	err := service.DeleteStepPort.Delete(command.Uuid)
	if err != nil {
		return errors.New("INTERNAL")
	}
	return nil
}
