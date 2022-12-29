package Service

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
	Recorder Repository.Recorder
}

func (creator *Creator) Create(
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.ExecutionMode,
) (Task.Task, error) {

	var task, _ = Task.New(
		host,
		port,
		stepVos,
		communicationMode,
		executionMode,
	)
	fmt.Println("TaskCreatedEvent")
	creator.Recorder.Persist(task)
	return task, nil
}
