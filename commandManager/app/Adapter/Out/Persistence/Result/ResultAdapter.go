package Result

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	"gorm.io/gorm"
)

type Adapter struct {
	Orm *gorm.DB
}

func (adapter Adapter) Save(result ResultDomain.Result) error {
	var taskMysql Task.Task
	response := adapter.Orm.First(&taskMysql, "uuid = ?", result.TaskUuid)
	if response.Error != nil {
		fmt.Printf("tasks %v. \n", response.Error)
		return response.Error
	}
	var resultMysql = FromDomain(result)
	resultMysql.TaskID = taskMysql.ID

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
		domainResults = append(domainResults, ToDomain(task))
	}
	return domainResults, nil
}
