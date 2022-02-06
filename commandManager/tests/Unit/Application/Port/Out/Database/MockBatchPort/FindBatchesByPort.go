package MockBatchPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindByMock struct {
	Result []Entity.Batch
}

func (mock FindByMock) FindBy(conditions interface{}) []Entity.Batch {
	return mock.Result
}
