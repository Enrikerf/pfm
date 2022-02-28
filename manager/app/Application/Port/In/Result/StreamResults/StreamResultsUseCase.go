package StreamResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UseCase interface {
	Stream(query Query) (results []Entity.Result, finish bool)
}
