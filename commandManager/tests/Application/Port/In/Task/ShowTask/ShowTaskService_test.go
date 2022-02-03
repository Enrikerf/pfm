package ShowTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ShowTask"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/Enrikerf/pfm/commandManager/tests/Application/Port/Out/Database/Task"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindResponse TaskDomain.Task
	taskFindError    error
}{
	{TaskDomain.Task{}, errors.New("not found")},
	{TaskDomain.Task{
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
	}, nil},
}

func TestListTaskService(t *testing.T) {
	for _, test := range tests {
		service := ShowTask.Service{FindByPort: Task.MockFindPort{
			Result: test.taskFindResponse,
			Error:  test.taskFindError,
		}}
		query := ShowTask.Query{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		tasks, err := service.Show(query)
		if test.taskFindError != nil {
			assert.Error(t, err)
		} else {
			assert.Equal(t, test.taskFindResponse, tasks, "error retrieving tasks")
		}

	}
}
