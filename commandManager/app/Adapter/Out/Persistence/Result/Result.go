package Result

import (
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Result struct {
	ID        uint      `gorm:"primaryKey"`
	Uuid      uuid.UUID `gorm:"size:255;not null;unique" json:"uuid"`
	TaskID    uint      `gorm:"size:255;not null;unique" json:"task_id"`
	TaskUuid  uuid.UUID `gorm:"size:255;not null;unique" json:"task_uuid"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func FromDomain(result ResultDomain.Result) Result {
	resultPersistence := Result{}
	resultPersistence.Uuid = result.Uuid
	resultPersistence.TaskUuid = result.TaskUuid
	resultPersistence.Content = result.Content
	return resultPersistence
}

func ToDomain(result Result) ResultDomain.Result {
	resultDomain := ResultDomain.Result{}
	resultDomain.Uuid = result.Uuid
	resultDomain.TaskUuid = result.TaskUuid
	resultDomain.Content = result.Content
	return resultDomain
}

func (result *Result) FindAll(db *gorm.DB) error {
	var err error
	err = db.Find(&result).Error
	if err != nil {
		return err
	}
	return nil
}
