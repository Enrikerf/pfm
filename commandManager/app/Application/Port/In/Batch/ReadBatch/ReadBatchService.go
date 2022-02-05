package ReadBatch

import (
	"errors"
	ResultFindPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
)

type Service struct {
	FindByPort ResultFindPort.FindBatch
}

func (service Service) Read(query Query) (Result.Batch, error) {
	batch, err := service.FindByPort.Find(query.Uuid)
	if err != nil {
		return Result.Batch{}, errors.New("not found")
	}
	return batch, nil
}
