package Task

import "github.com/google/uuid"

type Task struct {
	Id      uuid.UUID
	Host    string
	Port    string
	Command string
	Mode    TaskModes
	Status  TaskStatus
}

func NewTask(host string, port string, command string, mode string) (Task, error) {
	task := Task{}
	task.Id = uuid.New()
	task.Host = host
	task.Port = port
	task.Command = command
	taskMode, err := GetTaskMode(mode)
	if err != nil {
		return Task{}, err
	}
	task.Mode = taskMode
	task.Status = Pending
	return task, nil
}
