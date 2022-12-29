package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/google/uuid"
)

type Task interface {
	GetUuid() uuid.UUID
	GetUuidString() string
	GetHost() Host.Vo
	GetPort() Port.Vo
	GetSteps() []Step.Step
	GetCommunicationMode() CommunicationMode.Mode
	GetExecutionMode() ExecutionMode.ExecutionMode
	GetStatus() Status.Status
}

type task struct {
	uuid              uuid.UUID
	host              Host.Vo
	port              Port.Vo
	steps             []Step.Step
	communicationMode CommunicationMode.Mode
	executionMode     ExecutionMode.ExecutionMode
	status            Status.Status
}

func New(
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.ExecutionMode,
) (Task, error) {
	task := &task{}
	task.uuid = uuid.New()
	for _, stepVo := range stepVos {
		step := Step.New(stepVo)
		task.steps = append(task.steps, step)
	}
	task.host = host
	task.port = port
	task.executionMode = executionMode
	task.communicationMode = communicationMode
	task.status = Status.Pending
	return task, nil
}

func (t task) GetUuid() uuid.UUID {
	return t.uuid
}

func (t task) GetUuidString() string {
	return t.uuid.String()
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

func (t task) GetExecutionMode() ExecutionMode.ExecutionMode {
	return t.executionMode
}

func (t task) GetStatus() Status.Status {
	return t.status
}
