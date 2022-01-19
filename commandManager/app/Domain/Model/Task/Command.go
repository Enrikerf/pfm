package Task

import "github.com/google/uuid"

type Command struct {
	Uuid     uuid.UUID
	TaskUuid uuid.UUID
	Name     string
}

func NewCommand(TaskUuid uuid.UUID, name string) (Command, error) {
	result := Command{}
	result.Uuid = uuid.New()
	result.TaskUuid = TaskUuid
	result.Name = name
	return result, nil
}
