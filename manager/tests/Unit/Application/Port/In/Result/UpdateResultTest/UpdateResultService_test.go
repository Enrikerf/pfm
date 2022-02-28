package UpdateResultTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/UpdateResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockResultPort"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
factors && equivalence classes
	FindResultPort
		- result || error
	FindTaskPort
		- task  || error
	UpdateResultPort
		- nil   || error
tests
	case	FindResultPort		FindTaskPort		UpdateResultPort		result
	----------------------------------------------------------------------------
	1		err					err					err					error
	2		err					err 				nil					error
	3		err					task				err					error
	4		err 				task				nil					error
	5		result 				err					err					error
	6		result 				err 				nil					error
	7		result 				task				err					error
	8		result 				task				nil					error
*/
var tests = []struct {
	FindResultPort   ResultPort.Find
	FindTaskPort     TaskPort.Find
	UpdateResultPort ResultPort.Update
	command          UpdateResult.Command
	result           error
}{
	{
		FindResultPort: MockResultPort.FindMock{
			Result: Entity.Result{},
			Error:  errors.New(""),
		},
		FindTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New(""),
		},
		UpdateResultPort: MockResultPort.UpdateMock{
			Error: errors.New(""),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New(""),
	},
}

func TestCreateResultService(t *testing.T) {
	for _, test := range tests {
		service := UpdateResult.Service{
			FindResultPort:   test.FindResultPort,
			UpdateResultPort: test.UpdateResultPort,
		}
		err := service.Update(test.command)
		if test.result != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
