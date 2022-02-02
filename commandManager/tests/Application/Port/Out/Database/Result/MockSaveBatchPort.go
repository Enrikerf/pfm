package Result

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type MockSaveBatchPort struct {
	Result error
	Saves  Result.Batch
}

func (mock MockSaveBatchPort) Save(batch Result.Batch) error {
	mock.Saves = batch
	if mock.Result != nil {
		return mock.Result
	}
	return nil
}
