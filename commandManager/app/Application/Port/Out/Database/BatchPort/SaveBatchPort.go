package BatchPort

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type Save interface {
	Save(batch Result.Batch) error
}
