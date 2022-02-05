package ShowResult

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type UseCase interface {
	Show(query Query) (Result.Result, error)
}
