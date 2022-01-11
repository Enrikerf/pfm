package Call

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type RequestPort interface {
	Request(task Task.Task) Result.Result
}
