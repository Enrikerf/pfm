package Task

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindByPort interface {
	FindBy(conditions interface{}) ([]Task.Task, error)
}
