package Result

import (
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Result struct {
	ID uuid.UUID `gorm:"primary_key; unique; 
                      type:uuid; column:id; 
                      default:uuid_generate_v4()"`
	TaskID uuid.UUID `gorm:"primary_key; unique; 
                      type:uuid; column:task_id; 
                      default:uuid_generate_v4()"`
	Content    string    `gorm:"size:255;not null;unique" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func FromDomain(result ResultDomain.Result) Result {
	resultPersistence := Result{}
	resultPersistence.ID = result.Id
	resultPersistence.TaskID = result.TaskId
	resultPersistence.Content = result.Content
	return resultPersistence
}

func ToDomain(result Result) ResultDomain.Result {
	resultDomain := ResultDomain.Result{}
	resultDomain.Id = result.ID
	resultDomain.TaskId = result.TaskID
	resultDomain.Content = result.Content
	return resultDomain
}

func (result *Result) SaveResult(db *gorm.DB) error {
	var err error
	err = db.Create(&result).Error
	if err != nil {
		return err
	}
	return nil
}
func (result *Result) FindAll(db *gorm.DB) error {
	var err error
	err = db.Find(&result).Error
	if err != nil {
		return err
	}
	return nil
}
