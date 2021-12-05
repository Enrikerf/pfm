package Result

import "github.com/google/uuid"

type Result struct {
	Id      uuid.UUID
	TaskId  uuid.UUID
	Content string
}

func NewResult(taskId uuid.UUID, content string) (Result, error) {
	result := Result{}
	result.Id = uuid.New()
	result.TaskId = taskId
	result.Content = content
	return result, nil
}
