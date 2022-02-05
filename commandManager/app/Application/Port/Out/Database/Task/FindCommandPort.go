package Task

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type FindCommand interface {
	Find(uuid string) (Task.Command, error)
}
