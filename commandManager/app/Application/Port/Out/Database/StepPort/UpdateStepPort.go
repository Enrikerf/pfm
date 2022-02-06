package StepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type UpdateStepPort interface {
	Update(step Task.Step) error
}
