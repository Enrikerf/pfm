package MockStepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UpdateMock struct {
	Error error
}

func (mock UpdateMock) Update(step Entity.Step) error {
	return mock.Error
}
