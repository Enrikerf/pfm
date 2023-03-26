package ResultAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"gorm.io/gorm"
)

type PersistAdapter struct {
	Orm *gorm.DB
}

func (adapter PersistAdapter) Persist(result Result.Result) {
	var currentResultMysql Model.ResultDb
	var resultValuesToUpdate = Model.ResultDb{}
	resultValuesToUpdate.FromDomainV2(result)
	err := adapter.Orm.First(&currentResultMysql, "uuid = ?", resultValuesToUpdate.Uuid).Error
	if err != nil {
		_ = adapter.Orm.Create(&resultValuesToUpdate).Error
	}
	adapter.Orm.Model(&currentResultMysql).Updates(resultValuesToUpdate)
}
