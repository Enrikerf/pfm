package StepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type Save interface {
	Save(step Task.Step) error
}
