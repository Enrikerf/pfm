package CreateTask

import (
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct{
	savePort TaskOutPort.SavePort
}

func (service Service) Create(command Command) (TaskDomain.Task, error) {
	var task, _ = TaskDomain.NewTask(command.Host, command.Port, command.Command, command.Mode)
	newTask, err := service.savePort.Save(task)
	if err != nil {
		return TaskDomain.Task{}, err
	}
	return newTask, nil
}
