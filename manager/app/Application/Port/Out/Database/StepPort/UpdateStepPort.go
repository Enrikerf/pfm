package StepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Update interface {
	Update(step Entity.Step) error
}
