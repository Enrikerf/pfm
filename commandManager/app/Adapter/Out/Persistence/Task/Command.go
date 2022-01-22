package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"time"
)

type command struct {
	ID        uint      `gorm:"primaryKey"`
	Uuid      uuid.UUID `gorm:"size:255;not null;unique" json:"uuid"`
	TaskID    uint      `gorm:"size:255;not null;unique" json:"task_id"`
	TaskUuid  uuid.UUID `gorm:"size:255;not null;unique" json:"task_uuid"`
	Name      string    `json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (command *command) FromDomain(commandDomain TaskDomain.Command) {
	command.Uuid = commandDomain.Uuid
	command.TaskUuid = commandDomain.TaskUuid
	command.Name = commandDomain.Name
}

func (command *command) ToDomain() TaskDomain.Command {
	resultDomain := TaskDomain.Command{}
	resultDomain.Uuid = command.Uuid
	resultDomain.TaskUuid = command.TaskUuid
	resultDomain.Name = command.Name
	return resultDomain
}
