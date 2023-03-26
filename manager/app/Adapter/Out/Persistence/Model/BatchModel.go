package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/google/uuid"
	"time"
)

type BatchDb struct {
	ID        uint
	Uuid      uuid.UUID
	TaskID    uint
	TaskUuid  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BatchDb) TableName() string {
	return "batches"
}

func (batchModel *BatchDb) FromDomain(selfEntity Entity.Batch) {
	batchModel.Uuid = selfEntity.Uuid
	batchModel.TaskUuid = selfEntity.TaskUuid
	batchModel.CreatedAt = selfEntity.CreatedAt
	batchModel.UpdatedAt = selfEntity.UpdatedAt
	for _, resultDomain := range selfEntity.Results {
		result := ResultDb{}
		result.FromDomain(resultDomain)
	}
}

func (batchModel *BatchDb) ToDomain() Entity.Batch {
	selfEntity := Entity.Batch{}
	selfEntity.Uuid = batchModel.Uuid
	selfEntity.TaskUuid = batchModel.TaskUuid
	selfEntity.CreatedAt = batchModel.CreatedAt
	selfEntity.UpdatedAt = batchModel.UpdatedAt
	return selfEntity
}

func (batchModel *BatchDb) FromDomainV2(batch Result.Batch) {
	batchModel.Uuid = batch.GetId().GetUuid()
	batchModel.TaskUuid = batch.GetTaskId().GetUuid()
	batchModel.CreatedAt = batch.GetCreatedAt()
}

func (batchModel *BatchDb) ToDomainV2() (Result.Batch, error) {
	var results []Result.Result
	selfEntity := Result.LoadBatch(
		Result.LoadBatchId(batchModel.Uuid),
		Task.LoadId(batchModel.TaskUuid),
		results,
		batchModel.CreatedAt,
	)
	return selfEntity, nil
}
