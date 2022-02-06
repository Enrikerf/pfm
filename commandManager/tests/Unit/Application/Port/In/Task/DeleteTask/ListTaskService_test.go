package DeleteTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct{ taskFindError error }{
	{errors.New("internal")},
	{nil},
}

func TestListTaskService(t *testing.T) {
	for _, test := range tests {
		service := DeleteTask.Service{DeleteTaskPort: MockTaskPort.MockDeletePort{
			Error: test.taskFindError,
		}}
		command := DeleteTask.Command{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		err := service.Delete(command)
		if test.taskFindError != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
