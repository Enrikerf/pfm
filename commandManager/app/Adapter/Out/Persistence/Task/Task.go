package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID        uint      `gorm:"primaryKey"`
	Uuid      uuid.UUID `gorm:"size:255;not null;unique" json:"uuid"`
	Host      string    `gorm:"size:255;not null;unique" json:"host"`
	Port      string    `gorm:"size:255;not null;unique" json:"port"`
	Command   string    `gorm:"size:255;not null;unique" json:"command"`
	Mode      string    `gorm:"size:255;not null;unique" json:"mode"`
	Status    string    `gorm:"size:255;not null;unique" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func FromDomain(task TaskDomain.Task) Task {
	taskPersistence := Task{}
	taskPersistence.Uuid = task.Uuid
	taskPersistence.Host = task.Host
	taskPersistence.Port = task.Port
	taskPersistence.Command = task.Command
	taskPersistence.Mode = task.Mode.String()
	taskPersistence.Status = task.Status.String()
	return taskPersistence
}

func ToDomain(task Task) TaskDomain.Task {
	taskDomain := TaskDomain.Task{}
	taskDomain.Uuid = task.Uuid
	taskDomain.Host = task.Host
	taskDomain.Port = task.Port
	taskDomain.Command = task.Command
	taskDomain.Mode, _ = TaskDomain.GetTaskMode(task.Mode)
	taskDomain.Status, _ = TaskDomain.GetStatus(task.Status)
	return taskDomain
}
