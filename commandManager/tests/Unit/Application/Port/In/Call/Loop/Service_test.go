package Loop

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
	"testing"
)

var domainTasks = []Entity.Task{
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Steps: []Entity.Step{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Sentence: "ls",
		}},
		Mode:   ValueObject.Unary,
		Status: ValueObject.Pending,
	},
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Steps: []Entity.Step{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Sentence: "ls",
		}},
		Mode:   ValueObject.Unary,
		Status: ValueObject.Pending,
	},
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Steps: []Entity.Step{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Sentence: "ls",
		}},
		Mode:   ValueObject.Unary,
		Status: ValueObject.Pending,
	},
	{
		Uuid: uuid.UUID{},
		Host: "",
		Port: "",
		Steps: []Entity.Step{{
			Uuid:     uuid.UUID{},
			TaskUuid: uuid.UUID{},
			Sentence: "ls",
		}},
		Mode:   ValueObject.Unary,
		Status: ValueObject.Pending,
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
	taskFindByResult     []Entity.Task
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
	//for index, test := range tests {
	//	service := Loop.New(
	//		MockCallPort.MockRequest{},
	//		MockTaskPort.MockFindByPort{Result: test.taskFindByResult},
	//		MockTaskPort.MockUpdatePort{Return: test.taskUpdatePortResult},
	//		MockBatchPort.SaveMock{Error: nil},
	//		MockResultPort.MockSavePort{Result: test.resultSavePortResult},
	//	)
	//	fmt.Println(index)
	//	service.Iteration()
	//	if !test.finalResult {
	//		assert.True(t, service.GetExit())
	//	} else {
	//		for index := range domainTasks {
	//			assert.Equal(t, ValueObject.Done, domainTasks[index].Status)
	//		}
	//	}
	//	resetTasks()
	//}
	t.Skipf("pending of implementation")
}

func resetTasks() {
	//TODO: view problem with array of array references
	for index := range domainTasks {
		domainTasks[index].Status = ValueObject.Pending
	}
}
