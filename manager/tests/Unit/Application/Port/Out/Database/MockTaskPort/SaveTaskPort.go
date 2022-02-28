package MockTaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type SaveMock struct {
	Err error
}

func (mock SaveMock) Save(task Entity.Task) error {
	return mock.Err
}
