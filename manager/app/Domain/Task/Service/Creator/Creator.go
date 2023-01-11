package Creator

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type Creator struct {
	SaveRepository Repository.Save
}

func (creator *Creator) Create(
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
) Task.Task {

	var task = Task.New(
		host,
		port,
		stepVos,
		communicationMode,
		executionMode,
	)
	creator.SaveRepository.Persist(task)

	if task.GetExecutionMode() == ExecutionMode.Automatic {
		fmt.Println("TaskCreatedEvent: activate loop pending")
	}

	return task
}
