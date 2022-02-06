package ListSteps

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UseCase interface {
	List(query Query) []Entity.Step
}
