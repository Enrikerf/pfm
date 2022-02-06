package BatchPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Save interface {
	Save(batch Entity.Batch) error
}
