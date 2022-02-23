package Manual

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/google/uuid"
)

type UseCase interface {
	ExecuteTask(task *Entity.Task, batch uuid.UUID)
}
