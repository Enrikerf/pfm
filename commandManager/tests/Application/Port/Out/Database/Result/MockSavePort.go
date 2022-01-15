package Result

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"

type MockSavePort struct {
	Result error
	Saves  []Result.Result
}

func (mock MockSavePort) Save(result Result.Result) error {
	mock.Saves = append(mock.Saves, result)
	if mock.Result != nil {
		return mock.Result
	}
	return nil
}
