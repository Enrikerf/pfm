package CreateStepTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/CreateStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/StepPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockStepPort"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockTaskPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
factors & equivalence classes
	FindTaskPort
		- task || error
	SaveStepPort
		- nil || error
tests
	case	FindTaskPort	SaveStepPort		step
	----------------------------------------------------------
	1		err				err					error
	2		err				nil					done
	3		task			err					error
	4		task			nil					success
*/
var tests = []struct {
	findTaskPort TaskPort.Find
	saveStepPort StepPort.Save
	command      CreateStep.Command
	step         error
}{
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New("error"),
		},
		saveStepPort: MockStepPort.SaveMock{
			Error: errors.New("error"),
		},
		command: GetDefaultValidCommandMock(),
		step:    errors.New("error"),
	},
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  errors.New("error"),
		},
		saveStepPort: MockStepPort.SaveMock{
			Error: nil,
		},
		command: GetDefaultValidCommandMock(),
		step:    errors.New("error"),
	},
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  nil,
		},
		saveStepPort: MockStepPort.SaveMock{
			Error: errors.New("error"),
		},
		command: GetDefaultValidCommandMock(),
		step:    errors.New("error"),
	},
	{
		findTaskPort: MockTaskPort.MockFindPort{
			Result: Entity.Task{},
			Error:  nil,
		},
		saveStepPort: MockStepPort.SaveMock{
			Error: nil,
		},
		command: GetDefaultValidCommandMock(),
		step:    nil,
	},
}

func TestCreateStepService(t *testing.T) {
	for _, test := range tests {
		var service = CreateStep.Service{
			FindTaskPort: test.findTaskPort,
			SaveStepPort: test.saveStepPort,
		}
		_, err := service.Create(test.command)
		if test.step != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
