package MockStepPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindMock struct {
	Result Entity.Step
	Error  error
}

func (mock FindMock) Find(uuid string) (Entity.Step, error) {
	return mock.Result, mock.Error
}
