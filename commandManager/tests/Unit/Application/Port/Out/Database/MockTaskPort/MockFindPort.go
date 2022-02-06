package MockTaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type MockFindPort struct {
	Result Entity.Task
	Error  error
}

func (mock MockFindPort) Find(uuid string) (Entity.Task, error) {
	return mock.Result, mock.Error
}
