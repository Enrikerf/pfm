package StreamResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Call/Manual"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
)

type Service struct {
	FindBatchPort         BatchPort.Find
	FindTaskPort          TaskPort.Find
	UpdateTaskPort        TaskPort.Update
	FindResultsStreamPort ResultPort.FindStream
	ManualCallUseCase     Manual.UseCase
	started               bool
}

func (service *Service) Stream(query Query) (results []Entity.Result, finish bool) {
	//TODO: control not found
	batch, err := service.FindBatchPort.Find(query.BatchUuid.String())
	task, err := service.FindTaskPort.Find(batch.TaskUuid.String())
	if !service.started {
		task.Status = ValueObject.Running
		err := service.UpdateTaskPort.Update(task)
		if err != nil {
			return nil, true
		}
		go service.ManualCallUseCase.ExecuteTask(&task, query.BatchUuid)
		service.started = true
	}

	if err != nil {
		return nil, true
	}
	if task.Status != ValueObject.Running {
		service.started = false
		return nil, true
	}
	return service.FindResultsStreamPort.FindStream(query.BatchUuid.String(), query.LastDate), false
}
