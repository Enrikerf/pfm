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

func (taskPersistence *Task) FromDomain(task TaskDomain.Task) {
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
