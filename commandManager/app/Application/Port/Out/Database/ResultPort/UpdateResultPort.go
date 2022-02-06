package ResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Update interface {
	Update(result Result.Result) error
}
