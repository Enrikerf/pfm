package ResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Find interface {
	Find(uuid string) (Entity.Result, error)
}
