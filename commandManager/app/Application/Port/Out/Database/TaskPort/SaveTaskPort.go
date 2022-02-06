package TaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Save interface {
	Save(task Task.Task) error
}
