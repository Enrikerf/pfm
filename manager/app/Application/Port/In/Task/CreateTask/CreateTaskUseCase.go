package CreateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UseCase interface {
	Create(task Command) (Entity.Task, error)
}
