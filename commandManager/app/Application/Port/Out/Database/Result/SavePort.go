package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type SavePort interface {
	Save(result Result.Result) error
}
