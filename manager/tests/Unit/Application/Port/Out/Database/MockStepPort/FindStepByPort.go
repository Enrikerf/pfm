package MockStepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindByMock struct {
	Step []Entity.Step
}

func (mock FindByMock) FindBy(conditions interface{}) []Entity.Step {
	return mock.Step
}
