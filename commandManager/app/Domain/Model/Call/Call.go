package Call

import "github.com/google/uuid"

type Call struct {
	Id      uuid.UUID
	TaskId  uuid.UUID
	Content string
}

func NewCall(taskId uuid.UUID, content string) (Result, error) {
	result := Result{}
	result.Id = uuid.New()
	result.TaskId = taskId
	result.Content = content
	return result, nil
}
