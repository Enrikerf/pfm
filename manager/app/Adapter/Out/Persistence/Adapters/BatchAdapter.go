package Adapters

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"gorm.io/gorm"
)

type BatchAdapter struct {
	Orm *gorm.DB
}

func (adapter BatchAdapter) Find(uuid string) (ResultDomain.Batch, error) {
	var batchMysql = Model.Batch{}
	err := adapter.Orm.First(&batchMysql, "uuid = ?", uuid).Error
	if err != nil {
		return ResultDomain.Batch{}, err
	}

	return batchMysql.ToDomain(), nil
}

func (adapter BatchAdapter) Save(batch ResultDomain.Batch) error {
	var taskMysql Model.Task
	var batchMysql Model.Batch
	response := adapter.Orm.First(&taskMysql, "uuid = ?", batch.TaskUuid)
	if response.Error != nil {
		fmt.Printf("tasks %v.\n", response.Error)
		return response.Error
	}
	batchMysql.FromDomain(batch)
	batchMysql.TaskID = taskMysql.ID
	err := adapter.Orm.Create(&batchMysql).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter BatchAdapter) FindAll() ([]ResultDomain.Batch, error) {

	var results []Model.Batch
	var domainBatch []ResultDomain.Batch
	response := adapter.Orm.Find(&results)
	if response.Error != nil {
		fmt.Printf("tasks %v. \n", response.Error)
		return domainBatch, response.Error
	}
	for _, batch := range results {
		domainBatch = append(domainBatch, batch.ToDomain())
	}
	return domainBatch, nil
}

func (adapter BatchAdapter) Update(batch ResultDomain.Batch) error {
	var taskMysql Model.Task
	var currentBatchMysql Model.Batch
	var batchValuesToUpdate = Model.Batch{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", batch.TaskUuid).Error
	if err != nil {
		return err
	}
	err = adapter.Orm.First(&currentBatchMysql, "uuid = ?", batch.Uuid.String()).Error
	if err != nil {
		return err
	}
	batchValuesToUpdate.FromDomain(batch)
	batchValuesToUpdate.TaskID = taskMysql.ID
	adapter.Orm.Model(&currentBatchMysql).Updates(batchValuesToUpdate)
	return nil
}

func (adapter BatchAdapter) FindBy(conditions interface{}) []ResultDomain.Batch {

	var batches []Model.Batch
	domainBatches := []ResultDomain.Batch{}
	err := adapter.Orm.
		Table("batches").
		Where(conditions).
		Find(&batches).
		Error
	if err != nil {
		fmt.Printf("batches %v. \n", err)
		return nil
	}

	for _, batch := range batches {
		domainBatches = append(domainBatches, batch.ToDomain())
	}
	return domainBatches
}

func (adapter BatchAdapter) Delete(uuid string) error {
	var batchMysql = Model.Result{}
	err := adapter.Orm.Delete(&batchMysql, "uuid = ?", uuid).Error
	if err != nil {
		return err
	}
	return nil
}
