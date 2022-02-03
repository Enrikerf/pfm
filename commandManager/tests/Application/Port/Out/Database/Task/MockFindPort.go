package Task

import TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type MockFindPort struct {
	Result TaskDomain.Task
	Error  error
}

func (mock MockFindPort) Find(uuid string) (TaskDomain.Task, error) {
	return mock.Result, mock.Error
}
