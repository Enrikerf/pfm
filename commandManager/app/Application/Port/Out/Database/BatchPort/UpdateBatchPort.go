package BatchPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Update interface {
	Update(result Entity.Batch) error
}
