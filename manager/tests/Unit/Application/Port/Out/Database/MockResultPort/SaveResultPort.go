package MockResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type SaveMock struct {
	Err error
}

func (mock SaveMock) Save(result Entity.Result) error {
	return mock.Err
}
