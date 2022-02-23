package Adapters

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"gorm.io/gorm"
	"time"
)

type ResultAdapter struct {
	Orm *gorm.DB
}

func (adapter ResultAdapter) Save(result ResultDomain.Result) error {
	var batchMysql Model.Batch
	var resultMysql Model.Result
	response := adapter.Orm.First(&batchMysql, "uuid = ?", result.BatchUuid)
	if response.Error != nil {
		fmt.Printf("tasks %v. \n", response.Error)
		return response.Error
	}
	resultMysql.FromDomain(result)
	resultMysql.BatchID = batchMysql.ID

	err := adapter.Orm.Create(&resultMysql).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter ResultAdapter) FindBy(conditions interface{}) []ResultDomain.Result {

	var results []Model.Result
	domainResults := []ResultDomain.Result{}
	err := adapter.Orm.
		Table("results").
		Where(conditions).
		Find(&results).
		Error
	if err != nil {
		fmt.Printf("results %v. \n", err)
		return nil
	}

	for _, result := range results {
		domainResults = append(domainResults, result.ToDomain())
	}
	return domainResults
}

func (adapter ResultAdapter) FindStream(batchUuid string, lastDate time.Time) []ResultDomain.Result {

	var results []Model.Result
	domainResults := []ResultDomain.Result{}
	err := adapter.Orm.
		Table("results").
		Where("batch_uuid = ? AND created_at > ?", batchUuid, lastDate.String()).
		Order("created_at asc").
		Find(&results).
		Error
	if err != nil {
		fmt.Printf("results %v. \n", err)
		return nil
	}

	for _, result := range results {
		domainResults = append(domainResults, result.ToDomain())
	}
	return domainResults
}

func (adapter ResultAdapter) Delete(uuid string) error {
	var taskMysql = Model.Result{}
	err := adapter.Orm.Delete(&taskMysql, "uuid = ?", uuid).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter ResultAdapter) Update(result ResultDomain.Result) error {
	var batchMysql Model.Batch
	var currentResultMysql Model.Result
	var resultValuesToUpdate = Model.Result{}
	err := adapter.Orm.First(&batchMysql, "uuid = ?", result.BatchUuid).Error
	if err != nil {
		return err
	}
	err = adapter.Orm.First(&currentResultMysql, "uuid = ?", result.Uuid.String()).Error
	if err != nil {
		return err
	}
	resultValuesToUpdate.FromDomain(result)
	resultValuesToUpdate.BatchID = batchMysql.ID
	adapter.Orm.Model(&currentResultMysql).Updates(resultValuesToUpdate)
	return nil
}

func (adapter ResultAdapter) Find(uuid string) (ResultDomain.Result, error) {
	var resultMysql = Model.Result{}
	err := adapter.Orm.First(&resultMysql, "uuid = ?", uuid).Error
	if err != nil {
		return ResultDomain.Result{}, err
	}

	return resultMysql.ToDomain(), nil
}
