package ReadTask

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type UseCase interface {
	Read(query Query) (Task.Task, error)
}
