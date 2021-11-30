package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
)

type SavePort interface{
	Save(task Task.Task) (Task.Task,error)
}