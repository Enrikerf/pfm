package ResultAdapter

import (
	"fmt"
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
	err := adapter.Orm.First(&currentMysqlModel, "uuid = ?", modelToUpdate.Uuid).Error
	if err != nil {
		var taskDb Model.TaskDb
		err := adapter.Orm.First(&taskDb, "uuid = ?", batch.GetTaskId().GetUuidString()).Error
		if err != nil {
			fmt.Println("task not found")
		}
		modelToUpdate.FromDomainV2(batch)
		modelToUpdate.TaskID = taskDb.ID
		err = adapter.Orm.Create(&modelToUpdate).Error
		if err != nil {
			fmt.Println("batch can not be created")
		}
	} else {
		adapter.Orm.Model(&currentMysqlModel).Updates(modelToUpdate)
	}
}
