package Task

import "github.com/google/uuid"

type Task struct {
	Uuid          uuid.UUID
	Host          string
	Port          string
	Commands      []Command
	Mode          Modes
	Status        TaskStatus
	ExecutionMode ExecutionMode
}

func NewTask(host string, port string, commands []Command, mode string, executionMode string) (Task, error) {
	task := Task{}
	task.Uuid = uuid.New()
	task.Host = host
	task.Port = port
	task.Commands = commands
	taskMode, err := GetTaskMode(mode)
	if err != nil {
		return Task{}, err
	}
	taskExecutionMode, err := GetExecutionMode(executionMode)
	if err != nil {
		return Task{}, err
	}
	task.ExecutionMode = taskExecutionMode
	task.Mode = taskMode
	task.Status = Pending
	return task, nil
}
