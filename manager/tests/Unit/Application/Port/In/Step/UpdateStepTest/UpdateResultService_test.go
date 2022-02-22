package UpdateStepTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/UpdateStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockStepPort"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
factors && equivalence classes
	FindStepPort
		- step || error
	FindTaskPort
		- task  || error
	Update
		- nil   || error
tests
	case	FindStepPort		FindTaskPort		Update		step
	----------------------------------------------------------------------------
	1		err					err					err					error
	2		err					err 				nil					error
	3		err					task				err					error
	4		err 				task				nil					error
	5		step 				err					err					error
	6		step 				err 				nil					error
	7		step 				task				err					error
	8		step 				task				nil					error
*/
var tests = []struct {
	FindStepPort   StepPort.Find
	FindTaskPort   TaskPort.Find
	UpdateStepPort StepPort.Update
	command        UpdateStep.Command
	step           error
}{
	{
		FindStepPort: MockStepPort.FindMock{
			Result: Entity.Step{},
			Error:  errors.New(""),
		},
		FindTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New(""),
		},
		UpdateStepPort: MockStepPort.UpdateMock{
			Error: errors.New(""),
		},
		command: GetDefaultValidCommandMock(),
		step:    errors.New(""),
	},
}

func TestCreateStepService(t *testing.T) {
	for _, test := range tests {
		service := UpdateStep.Service{
			FindStepPort:   test.FindStepPort,
			FindTaskPort:   test.FindTaskPort,
			UpdateStepPort: test.UpdateStepPort,
		}
		err := service.Update(test.command)
		if test.step != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
