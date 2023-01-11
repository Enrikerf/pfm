package Eraser

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type Eraser struct {
	FindRepository   Repository.Find
	DeleteRepository Repository.Delete
}

func (eraser *Eraser) Erase(id Task.Id) error {

	task, err := eraser.FindRepository.Find(id)
	if err != nil {
		return NewTaskNotFoundError()
	}
	if task.GetStatus() == Status.Running {
		return NewRunningTaskCantBeDeletedError()
	}
	err = eraser.DeleteRepository.Delete(id)
	if err != nil {
		return NewAdapterError()
	}

	return nil
}
