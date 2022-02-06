package MockBatchPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UpdateMock struct {
	Err error
}

func (mock UpdateMock) Update(result Entity.Batch) error {
	return mock.Err
}
