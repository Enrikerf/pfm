package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type UpdateCommandPort interface {
	Update(task Task.Command) error
}
