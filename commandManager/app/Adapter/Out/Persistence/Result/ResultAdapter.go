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

func (adapter Adapter) FindAll() ([]ResultDomain.Result, error) {

	var results []Result
	var domainResults []ResultDomain.Result
	response := adapter.Orm.Find(&results)
	if response.Error != nil {
		fmt.Printf("tasks %v. \n", response.Error)
		return domainResults, response.Error
	}
	for _, task := range results {
		domainResults = append(domainResults, task.ToDomain())
	}
	return domainResults, nil
}
