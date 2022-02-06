package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type SaveTaskPort interface {
	Save(task Task.Task) error
}
