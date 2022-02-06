package TaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Update interface {
	Update(task Task.Task) error
}
