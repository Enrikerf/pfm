package CreateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Service struct {
	SaveTaskPort TaskPort.Save
}

func (service Service) Create(command Command) (TaskDomain.Task, error) {
	var commands []TaskDomain.Step
	var task, err = TaskDomain.NewTask(
		command.Host,
		command.Port,
		commands,
		command.Mode,
		command.ExecutionMode,
	)
	//TODO: this is responsibility of adapterIn not application
	//TODO: Vo needed because the dependency command-taskUuid taskVo has commandsVo without uuids
	for _, commandString := range command.Commands {
		newCommand, err := TaskDomain.NewCommand(task.Uuid, commandString)
		if err != nil {
			return TaskDomain.Task{}, err
		}
		task.Steps = append(task.Steps, newCommand)
	}
	if err != nil {
		return TaskDomain.Task{}, err
	}
	err = service.SaveTaskPort.Save(task)
	if err != nil {
		return TaskDomain.Task{}, err
	}
	return task, nil
}
