package MockTaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type MockFindByPort struct {
	Result []Entity.Task
}

func (mock MockFindByPort) FindBy(conditions interface{}) []Entity.Task {
	return mock.Result
}
