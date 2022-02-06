package Task

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindStep interface {
	Find(uuid string) (Task.Step, error)
}
