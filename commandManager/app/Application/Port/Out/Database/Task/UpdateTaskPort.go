package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type UpdateTaskPort interface {
	Update(task Task.Task) error
}
