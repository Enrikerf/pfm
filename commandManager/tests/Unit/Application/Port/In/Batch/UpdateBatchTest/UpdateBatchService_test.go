package UpdateBatchTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/UpdateBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockBatchPort"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
factors && equivalence classes
	FindBatchPort
		- batch || error
	FindTaskPort
		- task  || error
	UpdateBatchPort
		- nil   || error
tests
	case	FindBatchPort		FindTaskPort		UpdateBatchPort		result
	----------------------------------------------------------------------------
	1		err					err					err					error
	2		err					err 				nil					error
	3		err					task				err					error
	4		err 				task				nil					error
	5		batch 				err					err					error
	6		batch 				err 				nil					error
	7		batch 				task				err					error
	8		batch 				task				nil					error
*/
var tests = []struct {
	FindBatchPort   BatchPort.Find
	FindTaskPort    TaskPort.Find
	UpdateBatchPort BatchPort.Update
	command         UpdateBatch.Command
	result          error
}{
	{
		FindBatchPort: MockBatchPort.FindMock{
			Result: Entity.Batch{},
			Error:  errors.New(""),
		},
		FindTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New(""),
		},
		UpdateBatchPort: MockBatchPort.UpdateMock{
			Error: errors.New(""),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New(""),
	},
}

func TestCreateBatchService(t *testing.T) {
	for _, test := range tests {
		service := UpdateBatch.Service{
			FindBatchPort:   test.FindBatchPort,
			UpdateBatchPort: test.UpdateBatchPort,
		}
		err := service.Update(test.command)
		if test.result != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
