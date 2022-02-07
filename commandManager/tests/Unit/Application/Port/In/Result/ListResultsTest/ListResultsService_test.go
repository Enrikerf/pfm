package ListResultsTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ListResults"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockResultPort"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindByResult []Entity.Result
}{
	{nil},
	{[]Entity.Result{}},
	{[]Entity.Result{{
		Uuid:      uuid.UUID{},
		BatchUuid: uuid.UUID{},
		Content:   "",
	}}},
}

func TestListResultService(t *testing.T) {
	for _, test := range tests {
		service := ListResults.Service{FindResultsByPort: MockResultPort.FindByMock{Result: test.taskFindByResult}}
		query := ListResults.Query{}
		tasks := service.List(query)
		assert.Equal(t, test.taskFindByResult, tasks, "error retrieving tasks")
	}
}
