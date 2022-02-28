package MockResultPort

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
)

type FindByMock struct {
	Result []Entity.Result
}

func (mock FindByMock) FindBy(conditions interface{}) []Entity.Result {
	return mock.Result
}
