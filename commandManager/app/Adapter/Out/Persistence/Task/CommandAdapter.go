package Task

import (
	"fmt"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"gorm.io/gorm"
)

type CommandAdapter struct {
	Orm *gorm.DB
}

func (adapter CommandAdapter) Find(uuid string) (TaskDomain.Command, error) {
	var taskMysql = Command{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", uuid).Error
	if err != nil {
		return TaskDomain.Command{}, err
	}

	return taskMysql.ToDomain(), nil
}

func (adapter CommandAdapter) Save(commandDomain TaskDomain.Command) error {
	var taskMysql Task
	var commandMysql Command
	response := adapter.Orm.First(&taskMysql, "uuid = ?", commandDomain.TaskUuid)
	if response.Error != nil {
		fmt.Printf("tasks %v.\n", response.Error)
		return response.Error
	}
	commandMysql.FromDomain(commandDomain)
	commandMysql.TaskID = taskMysql.ID

	err := adapter.Orm.Create(&commandMysql).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter CommandAdapter) FindAll() ([]TaskDomain.Command, error) {

	var commands []Command
	var domainResults []TaskDomain.Command
	response := adapter.Orm.Find(&commands)
	if response.Error != nil {
		fmt.Printf("tasks %v. \n", response.Error)
		return domainResults, response.Error
	}
	for _, commandMysql := range commands {
		domainResults = append(domainResults, commandMysql.ToDomain())
	}
	return domainResults, nil
}

func (adapter CommandAdapter) Update(command TaskDomain.Command) error {
	var taskMysql Task
	var currentCommandMysql Command
	var commandValuesToUpdate = Command{}
	err := adapter.Orm.First(&taskMysql, "uuid = ?", command.TaskUuid).Error
	if err != nil {
		return err
	}
	err = adapter.Orm.First(&currentCommandMysql, "uuid = ?", command.Uuid.String()).Error
	if err != nil {
		return err
	}
	commandValuesToUpdate.FromDomain(command)
	commandValuesToUpdate.TaskID = taskMysql.ID
	adapter.Orm.Model(&currentCommandMysql).Updates(commandValuesToUpdate)
	return nil
}

func (adapter CommandAdapter) Delete(uuid string) error {
	var commandMysql = Command{}
	err := adapter.Orm.Delete(&commandMysql, "uuid = ?", uuid).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter CommandAdapter) FindBy(conditions interface{}) []TaskDomain.Command {

	var commands []Command
	domainCommands := []TaskDomain.Command{}
	err := adapter.Orm.
		Table("commands").
		Where(conditions).
		Find(&commands).
		Error
	if err != nil {
		fmt.Printf("commands %v. \n", err)
		return nil
	}

	for _, command := range commands {
		domainCommands = append(domainCommands, command.ToDomain())
	}
	return domainCommands
}
