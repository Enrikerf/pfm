package Entity

import (
	"github.com/google/uuid"
	"time"
)

type Batch struct {
	Uuid      uuid.UUID
	TaskUuid  uuid.UUID
	Results   []Result
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBatch(TaskUuid uuid.UUID, results []Result) Batch {
	batch := Batch{}
	batch.Uuid = uuid.New()
	batch.TaskUuid = TaskUuid
	batch.Results = results
	batch.CreatedAt = time.Now()
	batch.UpdatedAt = time.Now()
	return batch
}

func (batch *Batch) AddResult(result Result) {
	batch.Results = append(batch.Results, result)
}
