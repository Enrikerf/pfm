package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model"
)

type MockSaveBatchPort struct {
	Result error
	Saves  Entity.Batch
}

func (mock MockSaveBatchPort) Save(batch Entity.Batch) error {
	mock.Saves = batch
	if mock.Result != nil {
		return mock.Result
	}
	return nil
}
