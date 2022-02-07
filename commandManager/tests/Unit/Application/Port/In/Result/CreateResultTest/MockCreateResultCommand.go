package CreateResultTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/CreateResult"
)

const DEFAULT_UUID = "daa883d8-c9b0-4c49-9f9c-b87cd5603758"
const DEFAULT_CONTENT = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut ligula vitae"

func GetDefaultValidCommandMock() CreateResult.Command {
	return CreateResult.Command{
		Content:   DEFAULT_CONTENT,
		BatchUuid: DEFAULT_UUID,
	}
}
