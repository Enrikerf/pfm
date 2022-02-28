package ResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Save interface {
	Save(result Entity.Result) error
}
