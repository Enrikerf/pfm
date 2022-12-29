package CreateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type UseCase interface {
	Create(command Command) (Task.Task, error)
}

func New(recorder Repository.Recorder) UseCase {
	return &useCase{Service.Creator{Recorder: recorder}}
}

type useCase struct {
	creator Service.Creator
}

func (useCase *useCase) Create(command Command) (Task.Task, error) {
	host, err := Host.NewVo(command.Host)
	if err != nil {
		return nil, err
	}
	port, err := Port.NewVo(command.Port)
	if err != nil {
		return nil, err
	}
	communicationMode, err := CommunicationMode.FromString(command.Mode)
	if err != nil {
		return nil, err
	}
	executionMode, err := ExecutionMode.FromString(command.ExecutionMode)
	if err != nil {
		return nil, err
	}
	var stepVos []Step.Vo
	for _, commandSentence := range command.CommandSentences {
		stepVo, err := Step.NewVo(commandSentence)
		if err != nil {
			return nil, err
		}
		stepVos = append(stepVos, stepVo)
	}
	if err != nil {
		return nil, err
	}
	task, _ := useCase.creator.Create(
		host,
		port,
		stepVos,
		communicationMode,
		executionMode,
	)
	return task, nil
}
