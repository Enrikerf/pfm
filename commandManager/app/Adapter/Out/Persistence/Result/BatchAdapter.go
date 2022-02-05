package Result

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"gorm.io/gorm"
)

type BatchAdapter struct {
	Orm *gorm.DB
}

func (adapter BatchAdapter) Find(uuid string) (ResultDomain.Batch, error) {
	var batchMysql = Batch{}
	err := adapter.Orm.First(&batchMysql, "uuid = ?", uuid).Error
	if err != nil {
		return ResultDomain.Batch{}, err
	}

	return batchMysql.ToDomain(), nil
}

func (adapter BatchAdapter) Save(batch ResultDomain.Batch) error {
	var taskMysql Task.Task
	var batchMysql Batch
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

	var results []Batch
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
