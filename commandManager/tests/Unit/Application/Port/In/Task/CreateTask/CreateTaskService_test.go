package CreateTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct{ taskSaveError error }{
	{errors.New("internal")},
	{nil},
}

func TestCreateTaskService(t *testing.T) {
	for _, test := range tests {
		service := CreateTask.Service{
			SaveTaskPort: MockTaskPort.SaveMock{Err: test.taskSaveError},
		}
		command := CreateTask.Command{
			Host:             "",
			Port:             "",
			CommandSentences: nil,
			Mode:             "",
			Status:           "",
			ExecutionMode:    "",
		}
		_, err := service.Create(command)
		if test.taskSaveError != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
