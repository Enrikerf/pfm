package ShowBatchTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ReadBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockBatchPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindResponse Entity.Batch
	taskFindError    error
}{
	{Entity.Batch{}, errors.New("not found")},
	{Entity.Batch{}, nil},
}

func TestListBatchService(t *testing.T) {
	for _, test := range tests {
		service := ReadBatch.Service{FindBatchPort: MockBatchPort.FindMock{
			Result: test.taskFindResponse,
			Error:  test.taskFindError,
		}}
		query := ReadBatch.Query{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		tasks, err := service.Read(query)
		if test.taskFindError != nil {
			assert.Error(t, err)
		} else {
			assert.Equal(t, test.taskFindResponse, tasks, "error retrieving")
		}

	}
}
