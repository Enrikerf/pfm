package Task

import TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type MockUpdatePort struct {
	Return error
}

func (mock MockUpdatePort) Update(task TaskDomain.Task) error {
	return mock.Return
}
