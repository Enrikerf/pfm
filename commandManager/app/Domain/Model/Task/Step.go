package Task

import "github.com/google/uuid"

type Step struct {
	Uuid     uuid.UUID
	TaskUuid uuid.UUID
	Sentence string
}

func NewCommand(TaskUuid uuid.UUID, sentence string) (Step, error) {
	self := Step{}
	self.Uuid = uuid.New()
	self.TaskUuid = TaskUuid
	self.Sentence = sentence
	return self, nil
}
