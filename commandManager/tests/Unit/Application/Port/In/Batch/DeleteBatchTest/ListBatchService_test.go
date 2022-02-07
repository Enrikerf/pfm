package DeleteBatchTest

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/DeleteBatch"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Application/Port/Out/Database/MockBatchPort"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct{ deleteBatchPortResponse error }{
	{errors.New("internal")},
	{nil},
}

func TestListBatchService(t *testing.T) {
	for _, test := range tests {
		service := DeleteBatch.Service{DeleteBatchPort: MockBatchPort.DeleteMock{
			Error: test.deleteBatchPortResponse,
		}}
		command := DeleteBatch.Command{Uuid: "11a24a01-700d-4f55-8afe-a6ce3c461ead"}
		err := service.Delete(command)
		if test.deleteBatchPortResponse != nil {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
