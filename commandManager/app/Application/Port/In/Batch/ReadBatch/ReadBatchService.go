package ReadBatch

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindBatchPort BatchPort.Find
}

func (service Service) Read(query Query) (Result.Batch, error) {
	batch, err := service.FindBatchPort.Find(query.Uuid)
	if err != nil {
		return Result.Batch{}, errors.New("not found")
	}
	return batch, nil
}
