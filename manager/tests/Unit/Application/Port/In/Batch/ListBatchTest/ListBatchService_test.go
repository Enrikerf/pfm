package ListBatchTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ListBatches"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockBatchPort"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var tests = []struct {
	taskFindByResult []Entity.Batch
}{
	{nil},
	{[]Entity.Batch{}},
	{[]Entity.Batch{{
		Uuid:      uuid.UUID{},
		TaskUuid:  uuid.UUID{},
		Results:   nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}}},
}

func TestListBatchService(t *testing.T) {
	for _, test := range tests {
		service := ListBatches.Service{FindBatchesByPort: MockBatchPort.FindByMock{Result: test.taskFindByResult}}
		query := ListBatches.Query{}
		tasks := service.List(query)
		assert.Equal(t, test.taskFindByResult, tasks, "error retrieving tasks")
	}
}
