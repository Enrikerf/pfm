package CreateStep

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UseCase interface {
	Create(command Command) (Entity.Step, error)
}
