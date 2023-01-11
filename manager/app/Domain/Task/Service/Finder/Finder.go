package Finder

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

type Finder struct {
	FindRepository Repository.Find
}

func (finder *Finder) Find(id Task.Id) (Task.Task, error) {

	task, err := finder.FindRepository.Find(id)
	if err != nil {
		return nil, NewTaskNotFoundError()
	}
	return task, nil
}
