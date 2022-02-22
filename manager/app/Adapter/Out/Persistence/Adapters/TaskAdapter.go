package Adapters

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"gorm.io/gorm"
)

type TaskAdapter struct {
	Orm *gorm.DB
}

func (adapter TaskAdapter) Find(uuid string) (TaskDomain.Task, error) {
	var taskMysql = Model.Task{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", uuid).Error
	if err != nil {
		return TaskDomain.Task{}, err
	}

	return taskMysql.ToDomain(), nil
}

func (adapter TaskAdapter) Delete(uuid string) error {
	var taskMysql = Model.Task{}
	err := adapter.Orm.Delete(&taskMysql, "uuid = ?", uuid).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter TaskAdapter) Save(task TaskDomain.Task) error {
	var taskMysql = Model.Task{}
	taskMysql.FromDomain(task)
	err := adapter.Orm.Create(&taskMysql).Error
	if err != nil {
		return err
	}

	return nil
}

func (adapter TaskAdapter) Update(task TaskDomain.Task) error {
	var currentTaskMysql Model.Task
	var taskValuesToUpdate = Model.Task{}
	taskValuesToUpdate.FromDomain(task)
	err := adapter.Orm.First(&currentTaskMysql, "uuid = ?", taskValuesToUpdate.Uuid).Error
	if err != nil {
		return err
	}
	adapter.Orm.Model(&currentTaskMysql).Updates(taskValuesToUpdate)
	return nil
}

func (adapter TaskAdapter) FindBy(conditions interface{}) []TaskDomain.Task {

	var tasks []Model.Task
	domainTasks := []TaskDomain.Task{}
	err := adapter.Orm.
		Table("tasks").
		Preload("Steps").
		Where(conditions).
		Joins("inner join steps on steps.task_id = tasks.id").
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
