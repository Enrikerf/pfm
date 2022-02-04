package ListTasks

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"

type UseCase interface {
	List(query Query) []Task.Task
}
