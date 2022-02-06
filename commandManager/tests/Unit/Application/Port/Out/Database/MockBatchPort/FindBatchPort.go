package MockBatchPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindMock struct {
	Result Entity.Batch
	Err    error
}

func (mock FindMock) Find(uuid string) (Entity.Batch, error) {
	return mock.Result, mock.Err
}
