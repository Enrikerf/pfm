package MockTaskPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type MockUpdatePort struct {
	Return error
}

func (mock MockUpdatePort) Update(task Entity.Task) error {
	return mock.Return
}
