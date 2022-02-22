package ResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Update interface {
	Update(result Entity.Result) error
}
