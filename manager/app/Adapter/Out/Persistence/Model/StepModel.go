package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	StepDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/google/uuid"
	"time"
)

type StepDb struct {
	ID        uint
	Uuid      uuid.UUID
	TaskID    uint
	TaskUuid  uuid.UUID
	Sentence  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (StepDb) TableName() string {
	return "steps"
}

func (stepModel *StepDb) FromDomainV2(taskUuid uuid.UUID, selfEntity StepDomain.Step) {
	stepModel.Uuid = selfEntity.GetId().GetUuid()
	stepModel.TaskUuid = taskUuid
	stepModel.Sentence = selfEntity.GetSentence()
}

func (stepModel *StepDb) FromDomain(selfEntity Entity.Step) {
	stepModel.Uuid = selfEntity.Uuid
	stepModel.TaskUuid = selfEntity.TaskUuid
	stepModel.Sentence = selfEntity.Sentence
}

func (stepModel *StepDb) ToDomain() Entity.Step {
	selfEntity := Entity.Step{}
	selfEntity.Uuid = stepModel.Uuid
	selfEntity.TaskUuid = stepModel.TaskUuid
	selfEntity.Sentence = stepModel.Sentence
	return selfEntity
}

func (stepModel *StepDb) ToDomainV2() (StepDomain.Step, error) {

	vo, err := StepDomain.NewVo(stepModel.Sentence)
	if err != nil {
		return nil, err
	}
	selfEntity := StepDomain.Load(
		StepDomain.LoadId(stepModel.Uuid),
		vo,
	)

	return selfEntity, nil
}
