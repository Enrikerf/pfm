package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type Task interface {
	GetId() Id
	GetHost() Host.Vo
	GetPort() Port.Vo
	GetSteps() []Step.Step
	GetCommunicationMode() CommunicationMode.Mode
	GetExecutionMode() ExecutionMode.Mode
	GetStatus() Status.Status
}

type task struct {
	id                Id
	host              Host.Vo
	port              Port.Vo
	steps             []Step.Step
	communicationMode CommunicationMode.Mode
	executionMode     ExecutionMode.Mode
	status            Status.Status
}

func New(
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
) Task {
	task := &task{}
	task.id = NewId()
	for _, stepVo := range stepVos {
		step := Step.New(stepVo)
		task.steps = append(task.steps, step)
	}
	task.host = host
	task.port = port
	task.executionMode = executionMode
	task.communicationMode = communicationMode
	task.status = Status.Pending
	return task
}

func Load(
	id Id,
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
) Task {
	task := &task{}
	task.id = id
	for _, stepVo := range stepVos {
		step := Step.New(stepVo)
		task.steps = append(task.steps, step)
	}
	task.host = host
	task.port = port
	task.executionMode = executionMode
	task.communicationMode = communicationMode
	task.status = Status.Pending
	return task
}

func (t task) GetId() Id {
	return t.id
}

func (t task) GetHost() Host.Vo {
	return t.host
}

func (t task) GetPort() Port.Vo {
	return t.port
}

func (t task) GetSteps() []Step.Step {
	return t.steps
}

func (t task) GetCommunicationMode() CommunicationMode.Mode {
	return t.communicationMode
}

func (t task) GetExecutionMode() ExecutionMode.Mode {
	return t.executionMode
}

func (t task) GetStatus() Status.Status {
	return t.status
}
