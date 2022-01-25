package Call

import (
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
)

type MockCallRequestPort struct {
}

func (mock MockCallRequestPort) Request(task TaskDomain.Task) []ResultDomain.Result {
	return []ResultDomain.Result{
		{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Content:  "",
		},
	}
}
