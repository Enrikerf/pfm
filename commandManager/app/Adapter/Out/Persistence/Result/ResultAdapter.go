package Result

import (
	"fmt"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"gorm.io/gorm"
)

type Adapter struct {
	Orm *gorm.DB
}

func (adapter Adapter) Save(result ResultDomain.Result) error {
	var batchMysql Batch
	var resultMysql Result
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

func (adapter Adapter) FindBy(conditions interface{}) []ResultDomain.Result {

	var results []Result
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

func (adapter Adapter) Delete(uuid string) error {
	var taskMysql = Result{}
	err := adapter.Orm.Delete(&taskMysql, "uuid = ?", uuid).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter Adapter) Update(result ResultDomain.Result) error {
	var batchMysql Batch
	var currentResultMysql Result
	var resultValuesToUpdate = Result{}
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

func (adapter Adapter) Find(uuid string) (ResultDomain.Result, error) {
	var resultMysql = Result{}
	err := adapter.Orm.First(&resultMysql, "uuid = ?", uuid).Error
	if err != nil {
		return ResultDomain.Result{}, err
	}

	return resultMysql.ToDomain(), nil
}
