package ListStepsTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/ListSteps"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockStepPort"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	taskFindByStep []Entity.Step
}{
	{nil},
	{[]Entity.Step{}},
	{[]Entity.Step{{
		Uuid:     uuid.UUID{},
		TaskUuid: uuid.UUID{},
		Sentence: "",
	}}},
}

func TestListStepService(t *testing.T) {
	for _, test := range tests {
		service := ListSteps.Service{FindStepByPort: MockStepPort.FindByMock{Step: test.taskFindByStep}}
		query := ListSteps.Query{}
		tasks := service.List(query)
		assert.Equal(t, test.taskFindByStep, tasks, "error retrieving tasks")
	}
}
