package Result

import "github.com/google/uuid"

type Batch struct {
	Uuid     uuid.UUID
	TaskUuid uuid.UUID
	Results  []Result
}

func NewBatch(TaskUuid uuid.UUID, results []Result) Batch {
	batch := Batch{}
	batch.Uuid = uuid.New()
	batch.TaskUuid = TaskUuid
	batch.Results = results
	return batch
}

func (batch *Batch) AddResult(result Result) {
	batch.Results = append(batch.Results, result)
}
