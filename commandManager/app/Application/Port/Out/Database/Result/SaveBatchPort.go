package Result

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type SaveBatchPort interface {
	Save(batch Result.Batch) error
}
