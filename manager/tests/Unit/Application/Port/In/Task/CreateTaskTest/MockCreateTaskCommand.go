package CreateTaskTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

func GetDefaultValidCommandMock() CreateTask.Command {
	return CreateTask.Command{
		Host:             "",
		Port:             "",
		CommandSentences: nil,
		Mode:             ValueObject.Unary.String(),
		Status:           ValueObject.Pending.String(),
		ExecutionMode:    ValueObject.Automatic.String(),
	}
}

func GetDefaultInValidCommandMock() CreateTask.Command {
	return CreateTask.Command{
		Host:             "",
		Port:             "",
		CommandSentences: nil,
		Mode:             "",
		Status:           "",
		ExecutionMode:    "",
	}
}
