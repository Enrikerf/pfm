package Adapters

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"gorm.io/gorm"
)

type StepAdapter struct {
	Orm *gorm.DB
}

func (adapter StepAdapter) Save(selfDomain TaskDomain.Step) error {
	var taskMysql Model.Task
	var selfPersistence Model.Step
	response := adapter.Orm.First(&taskMysql, "uuid = ?", selfDomain.TaskUuid)
	if response.Error != nil {
		fmt.Printf("tasks %v.\n", response.Error)
		return response.Error
	}
	selfPersistence.FromDomain(selfDomain)
	selfPersistence.TaskID = taskMysql.ID

	err := adapter.Orm.Create(&selfPersistence).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter StepAdapter) Find(uuid string) (TaskDomain.Step, error) {
	var selfPersistence = Model.Step{}
	err := adapter.Orm.First(&selfPersistence, "uuid = ?", uuid).Error
	if err != nil {
		return TaskDomain.Step{}, err
	}

	return selfPersistence.ToDomain(), nil
}

func (adapter StepAdapter) Update(selfDomain TaskDomain.Step) error {
	var taskMysql Model.Task
	var selfPersistence Model.Step
	var selfValuesToUpdate = Model.Step{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", selfDomain.TaskUuid).Error
	if err != nil {
		return err
	}
	err = adapter.Orm.First(&selfPersistence, "uuid = ?", selfDomain.Uuid.String()).Error
	if err != nil {
		return err
	}
	selfValuesToUpdate.FromDomain(selfDomain)
	selfValuesToUpdate.TaskID = taskMysql.ID
	adapter.Orm.Model(&selfPersistence).Updates(selfValuesToUpdate)
	return nil
}

func (adapter StepAdapter) Delete(uuid string) error {
	var selfPersistence = Model.Step{}
	err := adapter.Orm.Delete(&selfPersistence, "uuid = ?", uuid).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter StepAdapter) FindBy(conditions interface{}) []TaskDomain.Step {

	var selfArrayPersistence []Model.Step
	selfArrayDomain := []TaskDomain.Step{}
	err := adapter.Orm.
		Table("steps").
		Where(conditions).
		Find(&selfArrayPersistence).
		Error
	if err != nil {
		fmt.Printf("steps %v. \n", err)
		return nil
	}

	for _, step := range selfArrayPersistence {
		selfArrayDomain = append(selfArrayDomain, step.ToDomain())
	}
	return selfArrayDomain
}
