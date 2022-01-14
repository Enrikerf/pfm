package Loop

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var domainTasks = []TaskDomain.Task{
	{
		Uuid:    uuid.UUID{},
		Host:    "",
		Port:    "",
		Command: "1000",
		Mode:    0,
		Status:  TaskDomain.Pending,
	},
	{
		Uuid:    uuid.UUID{},
		Host:    "",
		Port:    "",
		Command: "100",
		Mode:    0,
		Status:  TaskDomain.Pending,
	},
	{
		Uuid:    uuid.UUID{},
		Host:    "",
		Port:    "",
		Command: "10",
		Mode:    0,
		Status:  TaskDomain.Pending,
	},
	{
		Uuid:    uuid.UUID{},
		Host:    "",
		Port:    "",
		Command: "1",
		Mode:    0,
		Status:  TaskDomain.Pending,
	},
}

var results []Result.Result

func TestManager_Loop(t *testing.T) {
	service := Manager{
		MockCallRequestPort{},
		MockFindTasksByPort{},
		MockUpdateTaskPort{},
		MockSaveResultPort{},
	}
	service.Iteration()
	for index := range domainTasks {
		assert.Equal(t, TaskDomain.Done, domainTasks[index].Status)
		//assert.Equal(t, domainTasks[index].Command, results[index].Content)
	}
}

type MockCallRequestPort struct {
}

func (mock MockCallRequestPort) Request(task TaskDomain.Task) Result.Result {
	/*	sleepAmount, err := strconv.Atoi(task.Command)
		if err != nil {
			fmt.Printf("error: %v. \n", err)
		}
		time.Sleep(time.Duration(sleepAmount) * time.Second)*/
	return Result.Result{
		Uuid:     uuid.UUID{},
		TaskUuid: uuid.UUID{},
		Content:  task.Command,
	}
}

type MockFindTasksByPort struct {
}

func (mock MockFindTasksByPort) FindBy(conditions interface{}) ([]TaskDomain.Task, error) {

	return domainTasks, nil
}

type MockUpdateTaskPort struct {
}

func (mock MockUpdateTaskPort) Update(task TaskDomain.Task) error {
	return nil
}

type MockSaveResultPort struct {
}

func (mock MockSaveResultPort) Save(result Result.Result) error {
	results = append(results, result)
	return nil
}
