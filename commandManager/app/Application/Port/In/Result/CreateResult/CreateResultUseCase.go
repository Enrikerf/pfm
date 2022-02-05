package CreateResult

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type UseCase interface {
	Create(command Command) (Result.Result, error)
}
