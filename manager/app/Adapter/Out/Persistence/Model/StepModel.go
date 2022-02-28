package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
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

func (Step) TableName() string {
	return "steps"
}

func (stepModel *Step) FromDomain(selfEntity Entity.Step) {
	stepModel.Uuid = selfEntity.Uuid
	stepModel.TaskUuid = selfEntity.TaskUuid
	stepModel.Sentence = selfEntity.Sentence
}

func (stepModel *Step) ToDomain() Entity.Step {
	selfEntity := Entity.Step{}
	selfEntity.Uuid = stepModel.Uuid
	selfEntity.TaskUuid = stepModel.TaskUuid
	selfEntity.Sentence = stepModel.Sentence
	return selfEntity
}
