package StepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindBy interface {
	FindBy(conditions interface{}) []Entity.Step
}
