package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/google/uuid"
	"time"
)

type Batch struct {
	ID        uint
	Uuid      uuid.UUID
	TaskID    uint
	TaskUuid  uuid.UUID
	Results   []Result
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (batchModel *Batch) FromDomain(selfEntity Entity.Batch) {
	batchModel.Uuid = selfEntity.Uuid
	batchModel.TaskUuid = selfEntity.TaskUuid
	batchModel.CreatedAt = selfEntity.CreatedAt
	batchModel.UpdatedAt = selfEntity.UpdatedAt
	for _, resultDomain := range selfEntity.Results {
		result := Result{}
		result.FromDomain(resultDomain)
		batchModel.Results = append(batchModel.Results, result)
	}
}

func (batchModel *Batch) ToDomain() Entity.Batch {
	selfEntity := Entity.Batch{}
	selfEntity.Uuid = batchModel.Uuid
	selfEntity.TaskUuid = batchModel.TaskUuid
	selfEntity.CreatedAt = batchModel.CreatedAt
	selfEntity.UpdatedAt = batchModel.UpdatedAt
	for _, result := range batchModel.Results {
		selfEntity.AddResult(result.ToDomain())
	}
	return selfEntity
}
