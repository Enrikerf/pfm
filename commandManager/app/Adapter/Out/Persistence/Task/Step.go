package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"time"
)

type Step struct {
	ID        uint
	Uuid      uuid.UUID
	TaskID    uint
	TaskUuid  uuid.UUID
	Sentence  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (step *Step) FromDomain(selfDomain TaskDomain.Step) {
	step.Uuid = selfDomain.Uuid
	step.TaskUuid = selfDomain.TaskUuid
	step.Sentence = selfDomain.Sentence
}

func (step *Step) ToDomain() TaskDomain.Step {
	selfDomain := TaskDomain.Step{}
	selfDomain.Uuid = step.Uuid
	selfDomain.TaskUuid = step.TaskUuid
	selfDomain.Sentence = step.Sentence
	return selfDomain
}
