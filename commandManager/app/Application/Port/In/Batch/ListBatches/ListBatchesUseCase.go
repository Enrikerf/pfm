package ListBatches

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type UseCase interface {
	List(query Query) []Result.Batch
}
