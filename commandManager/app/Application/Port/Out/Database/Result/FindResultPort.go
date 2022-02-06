package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type FindResult interface {
	Find(uuid string) (Result.Result, error)
}
