package Entity

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
)

type Step struct {
	Uuid     uuid.UUID
	TaskUuid uuid.UUID
	Sentence string
}

func NewStep(TaskUuid uuid.UUID, stepVo ValueObject.StepVo) Step {
	self := Step{}
	self.Uuid = uuid.New()
	self.TaskUuid = TaskUuid
	self.Sentence = stepVo.Sentence
	return self
}
