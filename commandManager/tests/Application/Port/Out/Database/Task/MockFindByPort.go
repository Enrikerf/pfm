package Task

import TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type MockFindByPort struct {
	Result []TaskDomain.Task
}

func (mock MockFindByPort) FindBy(conditions interface{}) []TaskDomain.Task {
	return mock.Result
}
