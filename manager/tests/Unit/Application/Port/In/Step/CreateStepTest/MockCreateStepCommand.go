package CreateStepTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/CreateStep"
)

const DEFAULT_UUID = "daa883d8-c9b0-4c49-9f9c-b87cd5603758"
const DEFAULT_SENTENCE = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut ligula vitae"

func GetDefaultValidCommandMock() CreateStep.Command {
	return CreateStep.Command{
		TaskUuid: DEFAULT_UUID,
		Sentence: DEFAULT_SENTENCE,
	}
}
