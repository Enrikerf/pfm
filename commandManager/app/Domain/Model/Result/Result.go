package Result

import "github.com/google/uuid"

type Result struct {
	Uuid      uuid.UUID
	BatchUuid uuid.UUID
	Content   string
}

func NewResult(TaskUuid uuid.UUID, content string) (Result, error) {
	result := Result{}
	result.Uuid = uuid.New()
	result.BatchUuid = TaskUuid
	result.Content = content
	return result, nil
}
