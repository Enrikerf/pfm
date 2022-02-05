package ReadBatch

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type UseCase interface {
	Read(query Query) (Result.Batch, error)
}
