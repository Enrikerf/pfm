package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type UpdatePort interface {
	Update(result Result.Result) error
}
