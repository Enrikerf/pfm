package MockResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type UpdateMock struct {
	Error error
}

func (mock UpdateMock) Update(result Entity.Result) error {
	return mock.Error
}
