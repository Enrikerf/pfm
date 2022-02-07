package MockStepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type SaveMock struct {
	Error error
}

func (mock SaveMock) Save(step Entity.Step) error {
	return mock.Error
}
