package CreateTask

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTaskOutPort struct {
}

func (mock MockTaskOutPort)Save(task TaskDomain.Task) (TaskDomain.Task,error)  {
	return task,nil
}
func TestCreate(t *testing.T) {
	service := Service{MockTaskOutPort{}}
	command := Command{
		Host:    "Host",
		Port:    "Port",
		Command: "Command",
		Mode:    "Mode",
		Status:  "Status",
	}
	newTask, err := service.Create(command)

	assert.Nil(t, err, "shouldn't be error")
	assert.Equalf(t, TaskDomain.Pending, newTask.Status, "should initialise pending")
}
