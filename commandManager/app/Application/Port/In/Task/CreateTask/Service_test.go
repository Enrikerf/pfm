package CreateTask

import (
	"fmt"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTaskOutOkPort struct {
}

func (mock MockTaskOutOkPort) Save(task TaskDomain.Task) error {
	return nil
}

func TestCreateOk(t *testing.T) {
	service := Service{MockTaskOutOkPort{}}
	command := Command{
		Host:    "Host",
		Port:    "Port",
		Command: "Command",
		Mode:    "Unary",
	}
	newTask, err := service.Create(command)

	assert.Nil(t, err, "shouldn't be error")
	assert.Equalf(t, TaskDomain.Pending, newTask.Status, "should initialise pending")
}

type MockTaskOutErrorPort struct {
}

func (mock MockTaskOutErrorPort) Save(task TaskDomain.Task) error {
	return fmt.Errorf("error")
}

func TestCreateFail(t *testing.T) {
	service := Service{MockTaskOutErrorPort{}}
	command := Command{
		Host:    "Host",
		Port:    "Port",
		Command: "Command",
		Mode:    "Mode",
		Status:  "Status",
	}
	_, err := service.Create(command)
	assert.NotNil(t, err, "should be error")
}
