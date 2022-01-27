package Result

import (
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
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

func (result *Result) FromDomain(resultDomain ResultDomain.Result) {
	result.Uuid = resultDomain.Uuid
	result.BatchUuid = resultDomain.BatchUuid
	result.Content = resultDomain.Content
}

func (result *Result) ToDomain() ResultDomain.Result {
	resultDomain := ResultDomain.Result{}
	resultDomain.Uuid = result.Uuid
	resultDomain.BatchUuid = result.BatchUuid
	resultDomain.Content = result.Content
	return resultDomain
}
