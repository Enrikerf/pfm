package Loop

import (
	"errors"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Call/Loop"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	mockResultPort "github.com/Enrikerf/pfm/commandManager/tests/Application/Port/Out/Database/Result"
	mockTask "github.com/Enrikerf/pfm/commandManager/tests/Application/Port/Out/Database/Task"
	mockCall "github.com/Enrikerf/pfm/commandManager/tests/Application/Port/Out/Grcp/Call"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var domainTasks = []TaskDomain.Task{
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Commands: []TaskDomain.Command{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Name:     "ls",
		}},
		Mode:   TaskDomain.Unary,
		Status: TaskDomain.Pending,
	},
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Commands: []TaskDomain.Command{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Name:     "ls",
		}},
		Mode:   TaskDomain.Unary,
		Status: TaskDomain.Pending,
	},
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Commands: []TaskDomain.Command{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Name:     "ls",
		}},
		Mode:   TaskDomain.Unary,
		Status: TaskDomain.Pending,
	},
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Commands: []TaskDomain.Command{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Name:     "ls",
		}},
		Mode:   0,
		Status: TaskDomain.Pending,
	},
}

/**
factors
	taskFindByResult
	taskUpdatePortResult
	resultSavePortResult
equivalence classes
	taskFindByResult
		- array || nil(error)
	taskUpdatePortResult
		- error || nil
	resultSavePortResult
		- error || nil
tests
	case	resultSavePortResult	taskUpdatePortResult	taskFindByResult	result
	----------------------------------------------------------------------------------
	1		nil						nil						nil						error
	2		nil						nil						tasks					done
	3		nil						error					nil						error
	4		nil						error					tasks					error
	5		error					nil						nil						error
	6		error					nil						tasks					error
	7		error					error					nil						error
	8		error					error					tasks					error
*/
var tests = []struct {
	taskFindByResult     []TaskDomain.Task
	taskUpdatePortResult error
	resultSavePortResult error
	finalResult          bool
}{
	{
		nil,
		nil,
		nil,
		false,
	},
	{
		domainTasks,
		nil,
		nil,
		true,
	},
	{
		nil,
		errors.New("update error"),
		nil,
		false,
	},
	{
		domainTasks,
		errors.New("update error"),
		nil,
		false,
	},
	{
		nil,
		nil,
		errors.New("save error"),
		false,
	},
	{
		domainTasks,
		nil,
		errors.New("save error"),
		false,
	},
	{
		nil,
		errors.New("update error"),
		errors.New("save error"),
		false,
	},
	{
		domainTasks,
		errors.New("update error"),
		errors.New("save error"),
		false,
	},
}

func TestManager_Loop(t *testing.T) {
	for index, test := range tests {
		service := Loop.New(
			mockCall.MockCallRequestPort{},
			mockTask.MockFindByPort{Result: test.taskFindByResult},
			mockTask.MockUpdatePort{Return: test.taskUpdatePortResult},
			mockResultPort.MockSavePort{Result: test.resultSavePortResult},
		)
		fmt.Println(index)
		service.Iteration()
		if !test.finalResult {
			assert.True(t, service.GetExit())
		} else {
			for index := range domainTasks {
				assert.Equal(t, TaskDomain.Done, domainTasks[index].Status)
			}
		}
		resetTasks()
	}

}

func resetTasks() {
	//TODO: view problem with array of array references
	for index := range domainTasks {
		domainTasks[index].Status = TaskDomain.Pending
	}
}
