package Call

import (
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type MockCallRequestPort struct {
}

func (mock MockCallRequestPort) Request(task TaskDomain.Task) ResultDomain.Batch {
	return ResultDomain.Batch{}
}
