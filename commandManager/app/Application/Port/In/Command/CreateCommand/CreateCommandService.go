package CreateCommand

import (
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	FindTaskPort    TaskOutPort.Find
	SaveCommandPort TaskOutPort.SaveCommandPort
}

func (service Service) Create(command Command) (TaskDomain.Command, error) {
	task, err := service.FindTaskPort.Find(command.TaskUuid)
	if err != nil {
		return TaskDomain.Command{}, err
	}
	result, err := TaskDomain.NewCommand(task.Uuid, command.Name)
	if err != nil {
		return TaskDomain.Command{}, err
	}
	err = service.SaveCommandPort.Save(result)
	if err != nil {
		return TaskDomain.Command{}, err
	}
	return result, nil
}
