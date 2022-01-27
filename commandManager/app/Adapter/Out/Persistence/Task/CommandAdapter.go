package Task

import (
	"fmt"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"gorm.io/gorm"
)

type CommandAdapter struct {
	Orm *gorm.DB
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
