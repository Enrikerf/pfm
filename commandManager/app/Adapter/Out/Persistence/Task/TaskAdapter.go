package Task

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"gorm.io/gorm"
)

type Adapter struct {
	Orm *gorm.DB
}

func (adapter Adapter) Save(task TaskDomain.Task) error {
	var taskMysql = FromDomain(task)
	err := taskMysql.SaveUser(adapter.Orm)
	if err != nil {
		return err
	}
	return nil
}
