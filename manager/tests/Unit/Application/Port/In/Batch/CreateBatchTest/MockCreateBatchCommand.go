package CreateBatchTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/CreateBatch"
)

const DEFAULT_UUID = "daa883d8-c9b0-4c49-9f9c-b87cd5603758"

func GetDefaultValidCommandMock() CreateBatch.Command {
	return CreateBatch.Command{
		TaskUuid: DEFAULT_UUID,
	}
}
