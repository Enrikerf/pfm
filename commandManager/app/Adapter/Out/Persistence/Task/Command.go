package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"time"
)

type Command struct {
	ID        uint
	Uuid      uuid.UUID
	TaskID    uint
	TaskUuid  uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (command *Command) FromDomain(commandDomain TaskDomain.Command) {
	command.Uuid = commandDomain.Uuid
	command.TaskUuid = commandDomain.TaskUuid
	command.Name = commandDomain.Name
}

func (command *Command) ToDomain() TaskDomain.Command {
	resultDomain := TaskDomain.Command{}
	resultDomain.Uuid = command.Uuid
	resultDomain.TaskUuid = command.TaskUuid
	resultDomain.Name = command.Name
	return resultDomain
}
