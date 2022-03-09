package Entity

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"github.com/google/uuid"
	"time"
)

type Result struct {
	ID        uint
	Uuid      uuid.UUID
	BatchUuid uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewResult(BatchUuid uuid.UUID, content string) (Result, error) {
	result := Result{}
	result.Uuid = uuid.New()
	result.BatchUuid = BatchUuid
	result.Content = content
	result.CreatedAt = time.Now()
	result.UpdatedAt = time.Now()
	return result, nil
}

func FromResultVo(BatchUuid uuid.UUID, vo ValueObject.ResultVo) Result {
	result := Result{}
	result.Uuid = uuid.New()
	result.BatchUuid = BatchUuid
	result.Content = vo.Content
	result.CreatedAt = vo.CreatedAt
	result.UpdatedAt = vo.UpdatedAt
	return result
}
