package MockBatchPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type SaveMock struct {
	Err error
}

func (mock SaveMock) Save(batch Entity.Batch) error {
	return mock.Err
}
