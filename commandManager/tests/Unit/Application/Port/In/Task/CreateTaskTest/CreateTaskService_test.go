package CreateTaskTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskSaveError error
	command       CreateTask.Command
}{
	{
		taskSaveError: errors.New("error"),
		command:       GetDefaultInValidCommandMock(),
	},
	{
		taskSaveError: errors.New("error"),
		command:       GetDefaultValidCommandMock(),
	},
	{
		taskSaveError: errors.New("error"),
		command:       GetDefaultInValidCommandMock(),
	},
	{
		taskSaveError: nil,
		command:       GetDefaultValidCommandMock(),
	},
}

/**
factors
	saveTaskPort
	command
equivalence classes
	taskFindByResult
		- task || error
	command
		- invalid (host,port,command.Sentences) -> length, (mode,ExecutionMode) -> invalid
		- valid
tests
	case	taskFindByResult	command			result
	----------------------------------------------------------
	1		err						invalid		error
	2		err						valid		done
	3		task					invalid		error
	4		task					valid		success
*/
func TestCreateTaskService(t *testing.T) {
	for _, test := range tests {
		service := CreateTask.Service{
			SaveTaskPort: MockTaskPort.SaveMock{Err: test.taskSaveError},
		}
		_, err := service.Create(test.command)
		if test.taskSaveError != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
