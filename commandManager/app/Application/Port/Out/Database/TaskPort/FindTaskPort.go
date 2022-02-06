package TaskPort

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type Find interface {
	Find(uuid string) (Task.Task, error)
}
