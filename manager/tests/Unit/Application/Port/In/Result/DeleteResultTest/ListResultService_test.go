package DeleteResultTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/DeleteResult"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockResultPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct{ deleteResultPortResponse error }{
	{errors.New("internal")},
	{nil},
}

func TestListResultService(t *testing.T) {
	for _, test := range tests {
		service := DeleteResult.Service{DeleteResultPort: MockResultPort.DeleteMock{
			Error: test.deleteResultPortResponse,
		}}
		command := DeleteResult.Command{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		err := service.Delete(command)
		if test.deleteResultPortResponse != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
