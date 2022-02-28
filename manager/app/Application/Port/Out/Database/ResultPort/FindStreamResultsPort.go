package ResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"time"
)

type FindStream interface {
	FindStream(batchUuid string, lastDate time.Time) []Entity.Result
}
