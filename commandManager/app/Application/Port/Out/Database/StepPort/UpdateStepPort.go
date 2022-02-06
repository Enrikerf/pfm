package StepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UpdateStepPort interface {
	Update(step Entity.Step) error
}
