package CreateResultTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/CreateResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockBatchPort"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockResultPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
factors & equivalence classes
	FindBatchPort
		- batch || error
	SaveResultPort
		- nil || error
tests
	case	FindBatchPort	SaveResultPort		result
	----------------------------------------------------------
	1		err				err					error
	2		err				nil					done
	3		task			err					error
	4		task			nil					success
*/
var tests = []struct {
	findBatchPort  BatchPort.Find
	saveResultPort ResultPort.Save
	command        CreateResult.Command
	result         error
}{
	{
		findBatchPort: MockBatchPort.FindMock{
			Result: Entity.Batch{},
			Error:  errors.New("error"),
		},
		saveResultPort: MockResultPort.SaveMock{
			Err: errors.New("error"),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New("error"),
	},
	{
		findBatchPort: MockBatchPort.FindMock{
			Result: Entity.Batch{},
			Error:  errors.New("error"),
		},
		saveResultPort: MockResultPort.SaveMock{
			Err: nil,
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New("error"),
	},
	{
		findBatchPort: MockBatchPort.FindMock{
			Result: Entity.Batch{},
			Error:  nil,
		},
		saveResultPort: MockResultPort.SaveMock{
			Err: errors.New("error"),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New("error"),
	},
	{
		findBatchPort: MockBatchPort.FindMock{
			Result: Entity.Batch{},
			Error:  nil,
		},
		saveResultPort: MockResultPort.SaveMock{
			Err: nil,
		},
		command: GetDefaultValidCommandMock(),
		result:  nil,
	},
}

func TestCreateResultService(t *testing.T) {
	for _, test := range tests {
		var service = CreateResult.Service{
			FindBatchPort:  test.findBatchPort,
			SaveResultPort: test.saveResultPort,
		}
		_, err := service.Create(test.command)
		if test.result != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
