package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model"
)

type MockFindAllPort struct {
	Result []TaskDomain.Task
}

func (mock MockFindAllPort) FindAll(conditions interface{}) []TaskDomain.Task {
	return mock.Result
}
