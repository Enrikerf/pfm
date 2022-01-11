package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type UpdatePort interface {
	Update(task Task.Task) error
}
