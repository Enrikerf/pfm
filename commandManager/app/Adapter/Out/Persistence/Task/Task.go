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
	Steps         []Step
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
	for _, stepDomain := range task.Steps {
		step := Step{}
		step.FromDomain(stepDomain)
		taskPersistence.Steps = append(taskPersistence.Steps, step)
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
	for _, stepMysql := range taskPersistence.Steps {
		taskDomain.Steps = append(taskDomain.Steps, stepMysql.ToDomain())
	}
	taskDomain.Mode, _ = TaskDomain.GetTaskMode(taskPersistence.Mode)
	taskDomain.Status, _ = TaskDomain.GetStatus(taskPersistence.Status)
	taskDomain.ExecutionMode, _ = TaskDomain.GetExecutionMode(taskPersistence.ExecutionMode)
	return taskDomain
}
