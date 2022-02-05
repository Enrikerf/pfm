package Result

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type FindBatchesByPort interface {
	FindBy(conditions interface{}) []Result.Batch
}
