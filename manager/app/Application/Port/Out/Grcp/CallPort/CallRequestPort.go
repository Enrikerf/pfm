package CallPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type Request interface {
	Request(task Task.Task) Result.Batch
}
