package StepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type Save interface {
	Save(step Entity.Step) error
}
