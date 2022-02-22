package UpdateBatchTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/UpdateBatch"
)

const DEFAULT_UUID = "daa883d8-c9b0-4c49-9f9c-b87cd5603758"

func GetDefaultValidCommandMock() UpdateBatch.Command {
	return UpdateBatch.Command{
		Uuid: DEFAULT_UUID,
		TaskUuid: Model.OptionalString{
			Change: true,
			Value:  DEFAULT_UUID,
		},
	}
}
