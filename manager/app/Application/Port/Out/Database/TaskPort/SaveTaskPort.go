package TaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Save interface {
	Save(task Entity.Task) error
}
