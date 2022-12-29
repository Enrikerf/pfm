package Adapters

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"gorm.io/gorm"
)

type TaskAdapter struct {
	Orm *gorm.DB
}

func (adapter TaskAdapter) Find(uuid string) (TaskDomain.Task, error) {
	var taskMysql = Model.TaskDb{}
	var resultsMysql = []Model.StepDb{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", uuid).Error
	if err != nil {
		return TaskDomain.Task{}, err
	}
	err = adapter.Orm.
		Table("steps").
		Where("task_uuid = ?", taskMysql.Uuid).
		Find(&resultsMysql).
		Error
	if err != nil {
		fmt.Printf("results %v. \n", err)
		return TaskDomain.Task{}, err
	}
	taskMysql.Steps = resultsMysql

	return taskMysql.ToDomain(), nil
}

func (adapter TaskAdapter) Persist(task Task.Task) {
	var taskMysql = Model.TaskDb{}
	taskMysql.FromDomainV2(task)
	_ = adapter.Orm.Create(&taskMysql).Error
}

func (adapter TaskAdapter) Delete(uuid string) error {
	var taskMysql = Model.TaskDb{}
	err := adapter.Orm.Delete(&taskMysql, "uuid = ?", uuid).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter TaskAdapter) Update(task TaskDomain.Task) error {
	var currentTaskMysql Model.TaskDb
	var taskValuesToUpdate = Model.TaskDb{}
	taskValuesToUpdate.FromDomain(task)
	err := adapter.Orm.First(&currentTaskMysql, "uuid = ?", taskValuesToUpdate.Uuid).Error
	if err != nil {
		return err
	}
	adapter.Orm.Model(&currentTaskMysql).Updates(taskValuesToUpdate)
	return nil
}

func (adapter TaskAdapter) FindBy(conditions interface{}) []TaskDomain.Task {

	var tasks []Model.TaskDb
	domainTasks := []TaskDomain.Task{}
	err := adapter.Orm.
		Table("tasks").
		Preload("Steps").
		Where(conditions).
		Joins("left join steps on steps.task_id = tasks.id").
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
