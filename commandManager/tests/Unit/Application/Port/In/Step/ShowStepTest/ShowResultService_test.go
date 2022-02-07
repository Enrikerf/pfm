package ShowStepTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/ReadStep"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockStepPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindResponse Entity.Step
	taskFindError    error
}{
	{Entity.Step{}, errors.New("not found")},
	{Entity.Step{}, nil},
}

func TestListStepService(t *testing.T) {
	for _, test := range tests {
		service := ReadStep.Service{FindStepPort: MockStepPort.FindMock{
			Result: test.taskFindResponse,
			Error:  test.taskFindError,
		}}
		query := ReadStep.Query{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		tasks, err := service.Read(query)
		if test.taskFindError != nil {
			assert.Error(t, err)
		} else {
			assert.Equal(t, test.taskFindResponse, tasks, "error retrieving")
		}

	}
}
