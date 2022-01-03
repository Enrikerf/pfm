package Task

import (
	"fmt"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"gorm.io/gorm"
)

type Adapter struct {
	Orm *gorm.DB
}

func (adapter Adapter) Save(task TaskDomain.Task) error {
	var taskMysql = FromDomain(task)
	err := taskMysql.Save(adapter.Orm)
	if err != nil {
		return err
	}
	return nil
}

func (adapter Adapter) FindAll() ([]TaskDomain.Task,error) {

	var tasks []Task
	var domainTasks []TaskDomain.Task
	response := adapter.Orm.Find(&tasks)
	if response.Error != nil {
		fmt.Printf("tasks %v. \n", response.Error)
		return domainTasks,response.Error
	}
	for _, task := range tasks {
		domainTasks = append(domainTasks, ToDomain(task))
	}
	return domainTasks,nil
}
