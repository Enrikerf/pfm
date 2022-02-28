package MockResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindMock struct {
	Result Entity.Result
	Error  error
}

func (mock FindMock) Find(uuid string) (Entity.Result, error) {
	return mock.Result, mock.Error
}
