package Adapters

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Exception"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"gorm.io/gorm"
)

type TaskAdapter struct {
	Orm *gorm.DB
}

func (adapter TaskAdapter) Find(id Task.Id) (Task.Task, error) {
	var taskMysql = Model.TaskDb{}
	var stepsMysql = []Model.StepDb{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return nil, Exception.NewRepositoryException()
	}
	err = adapter.Orm.
		Table("steps").
		Where("task_uuid = ?", taskMysql.Uuid).
		Find(&stepsMysql).
		Error
	if err != nil {
		return nil, Finder.NewTaskNotFoundError()
	}
	taskMysql.Steps = stepsMysql

	task, err := taskMysql.ToDomainV2()
	if err != nil {
		return nil, Exception.NewRepositoryException()
	}
	return task, nil
}

func (adapter TaskAdapter) Persist(task Task.Task) {
	var taskMysql = Model.TaskDb{}
	taskMysql.FromDomainV2(task)
	_ = adapter.Orm.Create(&taskMysql).Error
	//TODO: go funcion to save this action in a queue and try it again
}

func (adapter TaskAdapter) Delete(id Task.Id) error {
	var taskMysql = Model.TaskDb{}
	err := adapter.Orm.Delete(&taskMysql, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return err
	}
	return nil
}
