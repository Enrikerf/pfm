package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/google/uuid"
	"time"
)

type ResultDb struct {
	ID        uint
	Uuid      uuid.UUID
	BatchID   uint
	BatchUuid uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ResultDb) TableName() string {
	return "results"
}

func (resultModel *ResultDb) FromDomain(selfEntity Entity.Result) {
	resultModel.Uuid = selfEntity.Uuid
	resultModel.BatchUuid = selfEntity.BatchUuid
	resultModel.Content = selfEntity.Content
}

func (resultModel *ResultDb) ToDomain() Entity.Result {
	selfEntity := Entity.Result{}
	selfEntity.ID = resultModel.ID
	selfEntity.Uuid = resultModel.Uuid
	selfEntity.BatchUuid = resultModel.BatchUuid
	selfEntity.Content = resultModel.Content
	selfEntity.CreatedAt = resultModel.CreatedAt
	selfEntity.UpdatedAt = resultModel.UpdatedAt
	return selfEntity
}

func (resultModel *ResultDb) FromDomainV2(result Result.Result) {
	selfEntity := Entity.Result{}
	selfEntity.Uuid = result.GetId().GetUuid()
	selfEntity.BatchUuid = result.GetBatchId().GetUuid()
	selfEntity.Content = result.GetContent().GetValue()
	selfEntity.CreatedAt = result.GetCreateAt()
}

func (resultModel *ResultDb) ToDomainV2() (Result.Result, error) {
	content, err := Content.NewContent(resultModel.Content)
	if err != nil {
		return nil, Error.NewRepositoryError()
	}
	selfEntity := Result.Load(
		Result.LoadId(resultModel.Uuid),
		Result.LoadBatchId(resultModel.BatchUuid),
		content,
		resultModel.CreatedAt,
	)
	return selfEntity, nil
}
