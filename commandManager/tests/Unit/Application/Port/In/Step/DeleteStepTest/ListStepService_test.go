package DeleteStepTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/DeleteStep"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockStepPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct{ deleteStepPortResponse error }{
	{errors.New("internal")},
	{nil},
}

func TestListStepService(t *testing.T) {
	for _, test := range tests {
		service := DeleteStep.Service{DeleteStepPort: MockStepPort.DeleteMock{
			Error: test.deleteStepPortResponse,
		}}
		command := DeleteStep.Command{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		err := service.Delete(command)
		if test.deleteStepPortResponse != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
