package ListTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/Enrikerf/pfm/commandManager/tests/Application/Port/Out/Database/Task"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindByResult []TaskDomain.Task
}{
	{nil},
	{[]TaskDomain.Task{}},
	{[]TaskDomain.Task{{
		Uuid: uuid.UUID{},
		Host: "host",
		Port: "port",
		Commands: []TaskDomain.Command{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Name:     "name",
		}},
		Mode:          0,
		Status:        0,
		ExecutionMode: 0,
	}}},
}

func TestListTaskService(t *testing.T) {
	for _, test := range tests {
		service := ListTasks.Service{FindByPort: Task.MockFindByPort{Result: test.taskFindByResult}}
		query := ListTasks.Query{}
		tasks := service.List(query)
		assert.Equal(t, test.taskFindByResult, tasks, "error retrieving tasks")
	}
}
