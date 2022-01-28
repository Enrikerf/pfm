package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID            uint
	Uuid          uuid.UUID
	Host          string
	Port          string
	Commands      []Command
	Mode          string
	Status        string
	ExecutionMode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (taskPersistence *Task) FromDomain(task TaskDomain.Task) {
	taskPersistence.Uuid = task.Uuid
	taskPersistence.Host = task.Host
	taskPersistence.Port = task.Port
	for _, commandDomain := range task.Commands {
		command := Command{}
		command.FromDomain(commandDomain)
		taskPersistence.Commands = append(taskPersistence.Commands, command)
	}
	taskPersistence.Mode = task.Mode.String()
	taskPersistence.Status = task.Status.String()
	taskPersistence.ExecutionMode = task.ExecutionMode.String()
}

func (taskPersistence *Task) ToDomain() TaskDomain.Task {
	taskDomain := TaskDomain.Task{}
	taskDomain.Uuid = taskPersistence.Uuid
	taskDomain.Host = taskPersistence.Host
	taskDomain.Port = taskPersistence.Port
	for _, commandMysql := range taskPersistence.Commands {
		taskDomain.Commands = append(taskDomain.Commands, commandMysql.ToDomain())
	}
	taskDomain.Mode, _ = TaskDomain.GetTaskMode(taskPersistence.Mode)
	taskDomain.Status, _ = TaskDomain.GetStatus(taskPersistence.Status)
	taskDomain.ExecutionMode, _ = TaskDomain.GetExecutionMode(taskPersistence.ExecutionMode)
	return taskDomain
}
