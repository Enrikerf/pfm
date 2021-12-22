package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID uuid.UUID `gorm:"primary_key; unique; 
                      type:uuid; column:id; 
                      default:uuid_generate_v4()"`
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
	taskPersistence.ID = task.Id
	taskPersistence.Host = task.Host
	taskPersistence.Port = task.Port
	taskPersistence.Command = task.Command
	taskPersistence.Mode = task.Mode.String()
	taskPersistence.Status = task.Status.String()
	return taskPersistence
}

func ToDomain(task Task) TaskDomain.Task {
	taskDomain := TaskDomain.Task{}
	taskDomain.Id = task.ID
	taskDomain.Host = task.Host
	taskDomain.Port = task.Port
	taskDomain.Command = task.Command
	taskDomain.Mode,_ = TaskDomain.GetTaskMode(task.Mode)
	taskDomain.Status,_ = TaskDomain.GetStatus(task.Status)
	return taskDomain
}

func (task *Task) SaveUser(db *gorm.DB) error {
	var err error
	err = db.Create(&task).Error
	if err != nil {
		return err
	}
	return nil
}
func (task *Task) FindAll(db *gorm.DB) error {
	var err error
	err = db.Find(&task).Error
	if err != nil {
		return err
	}
	return nil
}
