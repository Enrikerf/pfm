package CallPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Request interface {
	Request(task Entity.Task) Entity.Batch
}
