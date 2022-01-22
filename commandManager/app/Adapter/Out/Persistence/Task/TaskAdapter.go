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
	var taskMysql = Task{}
	taskMysql.FromDomain(task)
	err := adapter.Orm.Create(&taskMysql).Error
	if err != nil {
		return err
	}

	return nil
}

func (adapter Adapter) Update(task TaskDomain.Task) error {
	var taskUpdatedMysql Task
	var taskValuesToUpdate = Task{}
	taskValuesToUpdate.FromDomain(task)
	err := adapter.Orm.First(&taskUpdatedMysql, "uuid = ?", taskValuesToUpdate.Uuid).Error
	if err != nil {
		return err
	}
	adapter.Orm.Model(&taskUpdatedMysql).Updates(taskValuesToUpdate)
	return nil
}

func (adapter Adapter) FindAll() ([]TaskDomain.Task, error) {

	var tasks []Task
	var domainTasks []TaskDomain.Task
	var err error
	err = adapter.Orm.Find(&tasks).Error
	if err != nil {
		fmt.Printf("tasks %v. \n", err)
		return domainTasks, err
	}
	for _, task := range tasks {
		domainTasks = append(domainTasks, task.ToDomain())
	}
	return domainTasks, nil
}

func (adapter Adapter) FindBy(conditions interface{}) []TaskDomain.Task {

	var tasks []Task
	domainTasks := []TaskDomain.Task{}
	err := adapter.Orm.
		Table("tasks").
		Preload("Commands").
		Where(conditions).
		Joins("inner join commands on commands.task_id = tasks.id").
		Group("tasks.id").
		Find(&tasks).
		Error
	if err != nil {
		fmt.Printf("tasks %v. \n", err)
		return nil
	}

	for _, task := range tasks {
		domainTasks = append(domainTasks, task.ToDomain())
	}
	return domainTasks
}
