package CreateTask

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type UseCase interface {
	Create(task Command) (Task.Task, error)
}
