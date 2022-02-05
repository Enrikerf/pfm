package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type FindBatch interface {
	Find(uuid string) (Result.Batch, error)
}
