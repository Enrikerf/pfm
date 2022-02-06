package MockCallPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type MockRequest struct {
}

func (mock MockRequest) Request(task Entity.Task) Entity.Batch {
	return Entity.Batch{}
}
