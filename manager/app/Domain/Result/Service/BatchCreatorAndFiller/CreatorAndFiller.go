package BatchCreatorAndFiller

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Error"
	ResultEvent "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type CreatorAndFiller interface {
	CreateAndFill(taskId Task.Id) (Result.Batch, error)
}

func New(
	dispatcher Event.Dispatcher,
	findRepository TaskRepository.Find,
	saveTaskRepository TaskRepository.Save,
	saveBatchRepository Repository.SaveBatch,
) CreatorAndFiller {
	return &executor{
		dispatcher,
		Finder.Finder{FindRepository: findRepository},
		saveTaskRepository,
		saveBatchRepository,
	}
}

type executor struct {
	dispatcher          Event.Dispatcher
	taskFinder          Finder.Finder
	saveTaskRepository  TaskRepository.Save
	saveBatchRepository Repository.SaveBatch
}

func (e *executor) CreateAndFill(taskId Task.Id) (Result.Batch, error) {
	task, err := e.taskFinder.Find(taskId)
	if err != nil {
		return nil, err
	}
	if task.GetExecutionMode() == ExecutionMode.Manual {
		return nil, Error.NewTaskNotManualCanNotBeExecutedManuallyError()
	}
	batch := Result.NewBatch(task.GetId())
	e.saveBatchRepository.Persist(batch)
	task.SetStatus(Status.New(Status.Running))
	e.saveTaskRepository.Persist(task)
	go e.dispatcher.Dispatch(ResultEvent.NewBatchCreatedToBeFilled(batch.GetId().GetUuidString()))
	return batch, nil
}
