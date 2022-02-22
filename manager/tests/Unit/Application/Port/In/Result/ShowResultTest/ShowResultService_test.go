package ShowResultTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ReadResult"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockResultPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindResponse Entity.Result
	taskFindError    error
}{
	{Entity.Result{}, errors.New("not found")},
	{Entity.Result{}, nil},
}

func TestListResultService(t *testing.T) {
	for _, test := range tests {
		service := ReadResult.Service{FindResultPort: MockResultPort.FindMock{
			Result: test.taskFindResponse,
			Error:  test.taskFindError,
		}}
		query := ReadResult.Query{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		tasks, err := service.Read(query)
		if test.taskFindError != nil {
			assert.Error(t, err)
		} else {
			assert.Equal(t, test.taskFindResponse, tasks, "error retrieving")
		}

	}
}
