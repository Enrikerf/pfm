package CreateTaskTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/UpdateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	FindTaskPort   TaskPort.Find
	UpdateTaskPort TaskPort.Update
	command        UpdateTask.Command
	result         error
}{
	{
		FindTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New(""),
		},
		UpdateTaskPort: MockTaskPort.MockUpdatePort{
			Return: errors.New(""),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New(""),
	},
	{
		FindTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New(""),
		},
		UpdateTaskPort: MockTaskPort.MockUpdatePort{
			Return: nil,
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New(""),
	},
	{
		FindTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  nil,
		},
		UpdateTaskPort: MockTaskPort.MockUpdatePort{
			Return: errors.New(""),
		},
		command: GetDefaultValidCommandMock(),
		result:  errors.New(""),
	},
	{
		FindTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  nil,
		},
		UpdateTaskPort: MockTaskPort.MockUpdatePort{
			Return: nil,
		},
		command: GetDefaultValidCommandMock(),
		result:  nil,
	},
}

/**
factors && equivalence classes
	FindTaskPort
		- task,nil || ...,error
	UpdateTaskPort
		- nil || error
	Command
		- valid || invalid
tests
	case	taskFindByResult	UpdateTaskPort		result
	----------------------------------------------------------
	1		err					err				error
	2		err					nil				error
	3		task				err				error
	4		task				nil				success
*/
func TestCreateTaskService(t *testing.T) {
	for _, test := range tests {
		service := UpdateTask.Service{
			FindTaskPort:   test.FindTaskPort,
			UpdateTaskPort: test.UpdateTaskPort,
		}
		err := service.Update(test.command)
		if test.result != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
