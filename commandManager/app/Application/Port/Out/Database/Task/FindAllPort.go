package Task

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindAllPort interface {
	FindAll() ([]Task.Task,error)
}
