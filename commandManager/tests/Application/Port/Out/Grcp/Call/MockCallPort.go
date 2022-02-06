package Call

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model"
)

type MockCallRequestPort struct {
}

func (mock MockCallRequestPort) Request(task TaskDomain.Task) TaskDomain.Batch {
	return TaskDomain.Batch{}
}
