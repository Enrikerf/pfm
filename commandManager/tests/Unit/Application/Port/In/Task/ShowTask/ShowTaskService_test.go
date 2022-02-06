package ShowTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ReadTask"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindResponse Entity.Task
	taskFindError    error
}{
	{Entity.Task{}, errors.New("not found")},
	{Entity.Task{
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
	}, nil},
}

func TestListTaskService(t *testing.T) {
	for _, test := range tests {
		service := ReadTask.Service{FindTaskPort: MockTaskPort.MockFindPort{
			Result: test.taskFindResponse,
			Error:  test.taskFindError,
		}}
		query := ReadTask.Query{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		tasks, err := service.Read(query)
		if test.taskFindError != nil {
			assert.Error(t, err)
		} else {
			assert.Equal(t, test.taskFindResponse, tasks, "error retrieving tasks")
		}

	}
}
