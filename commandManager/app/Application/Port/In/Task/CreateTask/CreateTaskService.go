package CreateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

type Service struct {
	SaveTaskPort TaskPort.Save
}

func (service Service) Create(command Command) (Entity.Task, error) {
	var stepVos []ValueObject.StepVo
	for _, commandSentence := range command.CommandSentences {
		stepVo, err := ValueObject.NewStepVo(commandSentence)
		if err != nil {
			return Entity.Task{}, err
		}
		stepVos = append(stepVos, stepVo)
	}
	var task, err = Entity.NewTask(
		command.Host,
		command.Port,
		stepVos,
		command.Mode,
		command.ExecutionMode,
	)
	//TODO: this is responsibility of adapterIn not application
	//TODO: Vo needed because the dependency command-taskUuid taskVo has commandsVo without uuids

	if err != nil {
		return Entity.Task{}, err
	}
	err = service.SaveTaskPort.Save(task)
	if err != nil {
		return Entity.Task{}, err
	}
	return task, nil
}
