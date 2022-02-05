package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type UpdateBatchPort interface {
	Update(result Result.Batch) error
}
