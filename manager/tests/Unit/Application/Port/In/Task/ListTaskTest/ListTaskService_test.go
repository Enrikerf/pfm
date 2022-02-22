package ListTaskTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindByResult []Entity.Task
}{
	{nil},
	{[]Entity.Task{}},
	{[]Entity.Task{{
		Uuid: uuid.UUID{},
		Host: "host",
		Port: "port",
		Steps: []Entity.Step{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Sentence: "name",
		}},
		Mode:          0,
		Status:        0,
		ExecutionMode: 0,
	}}},
}

func TestListTaskService(t *testing.T) {
	for _, test := range tests {
		service := ListTasks.Service{FindTasksByPort: MockTaskPort.MockFindByPort{Result: test.taskFindByResult}}
		query := ListTasks.Query{}
		tasks := service.List(query)
		assert.Equal(t, test.taskFindByResult, tasks, "error retrieving tasks")
	}
}
