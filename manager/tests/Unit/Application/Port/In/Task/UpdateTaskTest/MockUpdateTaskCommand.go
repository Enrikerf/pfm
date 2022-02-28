package CreateTaskTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/UpdateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

const DEFAULT_UUID = "daa883d8-c9b0-4c49-9f9c-b87cd5603758"

func GetDefaultValidCommandMock() UpdateTask.Command {
	return UpdateTask.Command{
		Uuid: DEFAULT_UUID,
		Host: UpdateTask.OptionalString{
			Change: true,
			Value:  "newHost",
		},
		Port: UpdateTask.OptionalString{
			Change: false,
			Value:  "",
		},
		Mode: UpdateTask.OptionalString{
			Change: true,
			Value:  ValueObject.Unary.String(),
		},
		Status: UpdateTask.OptionalString{
			Change: false,
			Value:  "",
		},
		ExecutionMode: UpdateTask.OptionalString{
			Change: true,
			Value:  ValueObject.Automatic.String(),
		},
	}
}

func GetDefaultInValidCommandMock() UpdateTask.Command {
	return UpdateTask.Command{
		Uuid: DEFAULT_UUID,
		Host: UpdateTask.OptionalString{
			Change: true,
			Value:  "newHost",
		},
		Port: UpdateTask.OptionalString{
			Change: false,
			Value:  "",
		},
		Mode: UpdateTask.OptionalString{
			Change: true,
			Value:  "invalid mode",
		},
		Status: UpdateTask.OptionalString{
			Change: false,
			Value:  "",
		},
		ExecutionMode: UpdateTask.OptionalString{
			Change: true,
			Value:  "invalid execution mode",
		},
	}
}
