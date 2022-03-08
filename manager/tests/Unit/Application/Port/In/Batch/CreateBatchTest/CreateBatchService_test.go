package CreateBatchTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/CreateBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockBatchPort"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
factors & equivalence classes
	TaskPort.Find
		- task || error
	BatchPort.Save
		- nil || error
tests
	case	TaskPort.Find	BatchPort.Save		result
	----------------------------------------------------------
	1		err				err					error
	2		err				nil					done
	3		task			err					error
	4		task			nil					success
*/
var tests = []struct {
	findTaskPort  TaskPort.Find
	batchSavePort BatchPort.Save
	command       CreateBatch.Command
	result        error
}{
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New("error"),
		},
		batchSavePort: MockBatchPort.SaveMock{
			Err: errors.New("error"),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New("error"),
	},
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New("error"),
		},
		batchSavePort: MockBatchPort.SaveMock{
			Err: nil,
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New("error"),
	},
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  nil,
		},
		batchSavePort: MockBatchPort.SaveMock{
			Err: errors.New("error"),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New("error"),
	},
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  nil,
		},
		batchSavePort: MockBatchPort.SaveMock{
			Err: nil,
		},
		command: GetDefaultValidCommandMock(),
		result:  nil,
	},
}

func TestCreateBatchService(t *testing.T) {
	for _, test := range tests {
		var service = CreateBatch.Service{
			FindTaskPort:  test.findTaskPort,
			SaveBatchPort: test.batchSavePort,
		}
		_, err := service.Create(test.command)
		if test.result != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
