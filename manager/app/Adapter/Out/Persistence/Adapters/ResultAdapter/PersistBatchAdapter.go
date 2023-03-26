package ResultAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"gorm.io/gorm"
)

type PersistBatchAdapter struct {
	Orm *gorm.DB
}

func (adapter PersistBatchAdapter) Persist(batch Result.Batch) {
	var currentMysqlModel Model.BatchDb
	var modelToUpdate = Model.BatchDb{}
	modelToUpdate.FromDomainV2(batch)
	err := adapter.Orm.First(&currentMysqlModel, "uuid = ?", modelToUpdate.Uuid).Error
	if err != nil {
		_ = adapter.Orm.Create(&modelToUpdate).Error
	}
	adapter.Orm.Model(&currentMysqlModel).Updates(modelToUpdate)
}
