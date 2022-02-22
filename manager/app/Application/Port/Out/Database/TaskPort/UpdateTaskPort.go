package TaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Update interface {
	Update(task Entity.Task) error
}
