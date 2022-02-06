package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type SaveStepPort interface {
	Save(step Task.Step) error
}
