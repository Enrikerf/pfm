package TaskAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	Error2 "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Error"
	"gorm.io/gorm"
)

type FindAdapter struct {
	Orm *gorm.DB
}

func (adapter FindAdapter) Find(id Task.Id) (Task.Task, error) {
	var taskMysql = Model.TaskDb{}
	var stepsMysql = []Model.StepDb{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return nil, Error.NewRepositoryError()
	}
	err = adapter.Orm.
		Table("steps").
		Where("task_uuid = ?", taskMysql.Uuid).
		Find(&stepsMysql).
		Error
	if err != nil {
		return nil, Error2.NewTaskNotFoundError()
	}
	taskMysql.Steps = stepsMysql

	task, err := taskMysql.ToDomainV2()
	if err != nil {
		return nil, Error.NewRepositoryError()
	}
	return task, nil
}
