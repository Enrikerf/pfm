package Result

import (
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
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

func (batch *Batch) FromDomain(batchDomain ResultDomain.Batch) {
	batch.Uuid = batchDomain.Uuid
	batch.TaskUuid = batchDomain.TaskUuid
	for _, resultDomain := range batchDomain.Results {
		result := Result{}
		result.FromDomain(resultDomain)
		batch.Results = append(batch.Results, result)
	}
}

func (batch *Batch) ToDomain() ResultDomain.Batch {
	batchDomain := ResultDomain.Batch{}
	batchDomain.Uuid = batch.Uuid
	batchDomain.TaskUuid = batch.TaskUuid
	for _, result := range batch.Results {
		batchDomain.AddResult(result.ToDomain())
	}
	return batchDomain
}
