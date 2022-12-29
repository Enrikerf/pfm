package Repository

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Task"

type Recorder interface {
	Persist(task Task.Task)
}
