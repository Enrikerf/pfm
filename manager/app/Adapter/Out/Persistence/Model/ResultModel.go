package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/google/uuid"
	"time"
)

type Result struct {
	ID        uint
	Uuid      uuid.UUID
	BatchID   uint
	BatchUuid uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (resultModel *Result) FromDomain(selfEntity Entity.Result) {
	resultModel.Uuid = selfEntity.Uuid
	resultModel.BatchUuid = selfEntity.BatchUuid
	resultModel.Content = selfEntity.Content
}

func (resultModel *Result) ToDomain() Entity.Result {
	selfEntity := Entity.Result{}
	selfEntity.ID = resultModel.ID
	selfEntity.Uuid = resultModel.Uuid
	selfEntity.BatchUuid = resultModel.BatchUuid
	selfEntity.Content = resultModel.Content
	selfEntity.CreatedAt = resultModel.CreatedAt
	selfEntity.UpdatedAt = resultModel.UpdatedAt
	return selfEntity
}
