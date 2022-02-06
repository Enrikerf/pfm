package MockResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type MockSavePort struct {
	Result error
	Saves  []Entity.Result
}

func (mock MockSavePort) Save(result Entity.Result) error {
	mock.Saves = append(mock.Saves, result)
	if mock.Result != nil {
		return mock.Result
	}
	return nil
}
