package ResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindStream interface {
	FindStream(batchUuid string, lastDate uint) []Entity.Result
}
