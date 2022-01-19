package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID            uint      `gorm:"primaryKey"`
	Uuid          uuid.UUID `gorm:"size:255;not null;unique" json:"uuid"`
	Host          string    `gorm:"size:255;not null;unique" json:"host"`
	Port          string    `gorm:"size:255;not null;unique" json:"port"`
	Commands      []command `gorm:"size:255;not null;unique" json:"command"`
	Mode          string    `gorm:"size:255;not null;unique" json:"mode"`
	Status        string    `gorm:"size:255;not null;unique" json:"status"`
	ExecutionMode string    `gorm:"column:execution_mode;size:255;not null;unique" json:"execution_mode"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (taskPersistence Task) FromDomain(task TaskDomain.Task) {
	taskPersistence.Uuid = task.Uuid
	taskPersistence.Host = task.Host
	taskPersistence.Port = task.Port
	for _, commandDomain := range task.Commands {
		command := command{}
		command.FromDomain(commandDomain)
		taskPersistence.Commands = append(taskPersistence.Commands, command)
	}
	taskPersistence.Mode = task.Mode.String()
	taskPersistence.Status = task.Status.String()
	taskPersistence.ExecutionMode = task.ExecutionMode.String()
}

func (taskPersistence Task) ToDomain(task Task) TaskDomain.Task {
	taskDomain := TaskDomain.Task{}
	taskDomain.Uuid = task.Uuid
	taskDomain.Host = task.Host
	taskDomain.Port = task.Port
	for _, commandMysql := range task.Commands {
		taskDomain.Commands = append(taskDomain.Commands, command{}.ToDomain(commandMysql))
	}
	taskDomain.Mode, _ = TaskDomain.GetTaskMode(task.Mode)
	taskDomain.Status, _ = TaskDomain.GetStatus(task.Status)
	taskDomain.ExecutionMode, _ = TaskDomain.GetExecutionMode(task.ExecutionMode)
	return taskDomain
}
