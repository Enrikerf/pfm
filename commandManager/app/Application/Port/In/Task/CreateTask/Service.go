package CreateTask

import (
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	SavePort TaskOutPort.SavePort
}

func (service Service) Create(command Command) (TaskDomain.Task, error) {
	var commands []TaskDomain.Command
	var task, err = TaskDomain.NewTask(command.Host, command.Port, commands, command.Mode, command.ExecutionMode)
	for _, commandString := range command.Commands {
		newCommand, err := TaskDomain.NewCommand(task.Uuid, commandString)
		if err != nil {
			return TaskDomain.Task{}, err
		}
		task.Commands = append(task.Commands, newCommand)
	}
	if err != nil {
		return TaskDomain.Task{}, err
	}
	err = service.SavePort.Save(task)
	if err != nil {
		return TaskDomain.Task{}, err
	}
	return task, nil
}
