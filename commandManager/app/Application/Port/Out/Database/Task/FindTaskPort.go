package Task

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindTask interface {
	Find(uuid string) (Task.Task, error)
}
