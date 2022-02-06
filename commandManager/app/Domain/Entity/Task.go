package Entity

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
)

type Task struct {
	Uuid          uuid.UUID
	Host          string
	Port          string
	Steps         []Step
	Mode          ValueObject.Modes
	Status        ValueObject.TaskStatus
	ExecutionMode ValueObject.ExecutionMode
}

func NewTask(host string, port string, stepVos []ValueObject.StepVo, mode string, executionMode string) (Task, error) {
	task := Task{}
	task.Uuid = uuid.New()
	for _, stepVo := range stepVos {
		step := NewStep(task.Uuid, stepVo)
		task.Steps = append(task.Steps, step)
	}
	task.Host = host
	task.Port = port
	//TODO: taskMode, ExecutionMode are Vos but how to handle errors? make strings not iota?
	taskMode, err := ValueObject.GetTaskMode(mode)
	if err != nil {
		return Task{}, err
	}
	taskExecutionMode, err := ValueObject.GetExecutionMode(executionMode)
	if err != nil {
		return Task{}, err
	}
	task.ExecutionMode = taskExecutionMode
	task.Mode = taskMode
	task.Status = ValueObject.Pending
	return task, nil
}
