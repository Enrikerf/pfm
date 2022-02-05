package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Find interface {
	Find(uuid string) (Result.Result, error)
}
