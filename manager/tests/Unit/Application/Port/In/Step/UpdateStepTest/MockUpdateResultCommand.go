package UpdateStepTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/UpdateStep"
)

const DEFAULT_UUID = "daa883d8-c9b0-4c49-9f9c-b87cd5603758"
const DEFAULT_CONTENT = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut ligula vitae"

func GetDefaultValidCommandMock() UpdateStep.Command {
	return UpdateStep.Command{
		Uuid: DEFAULT_UUID,
		TaskUuid: Model.OptionalString{
			Change: true,
			Value:  DEFAULT_CONTENT,
		},
		Sentence: Model.OptionalString{
			Change: true,
			Value:  DEFAULT_CONTENT,
		},
	}
}
