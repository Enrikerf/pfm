package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type UpdateResultPort interface {
	Update(result Result.Result) error
}
