package ReadStep

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UseCase interface {
	Read(query Query) (Entity.Step, error)
}
