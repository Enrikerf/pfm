package UpdateResultTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/UpdateResult"
)

const DEFAULT_UUID = "daa883d8-c9b0-4c49-9f9c-b87cd5603758"
const DEFAULT_CONTENT = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut ligula vitae"

func GetDefaultValidCommandMock() UpdateResult.Command {
	return UpdateResult.Command{
		Uuid: DEFAULT_UUID,
		Content: Model.OptionalString{
			Change: true,
			Value:  DEFAULT_CONTENT,
		},
		BatchUuid: Model.OptionalString{
			Change: true,
			Value:  DEFAULT_UUID,
		},
	}
}
