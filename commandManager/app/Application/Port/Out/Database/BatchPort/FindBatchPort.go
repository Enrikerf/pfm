package BatchPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Find interface {
	Find(uuid string) (Entity.Batch, error)
}
