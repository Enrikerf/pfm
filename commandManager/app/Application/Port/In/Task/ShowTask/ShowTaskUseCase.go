package ShowTask

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type UseCase interface {
	Show(query Query) (Task.Task, error)
}
