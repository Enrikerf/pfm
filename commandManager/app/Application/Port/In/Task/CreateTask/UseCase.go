package CreateTask

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type UseCase interface {
	Save(task Command) (Task.Task, error)
}
