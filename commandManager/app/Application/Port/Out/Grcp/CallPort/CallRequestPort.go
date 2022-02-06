package CallPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Request interface {
	Request(task Task.Task) Result.Batch
}
