package CreateBatchAndFill

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Service/BatchCreatorAndFiller"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

type UseCase interface {
	Execute(command Command) (Result.Batch, error)
}

func New(
	dispatcher Event.Dispatcher,
	findTaskRepository TaskRepository.Find,
	saveTaskRepository TaskRepository.Save,
	saveBatchRepository Repository.SaveBatch,
) UseCase {
	return &useCase{
		batchCreatorAndFiller: BatchCreatorAndFiller.New(
			dispatcher,
			findTaskRepository,
			saveTaskRepository,
			saveBatchRepository,
		),
	}
}

type useCase struct {
	batchCreatorAndFiller BatchCreatorAndFiller.CreatorAndFiller
}

func (service *useCase) Execute(command Command) (Result.Batch, error) {
	taskId, err := Task.LoadIdFromString(command.TaskUuid)
	if err != nil {
		return nil, err
	}
	batch, err := service.batchCreatorAndFiller.CreateAndFill(taskId)
	if err != nil {
		return nil, err
	}
	return batch, nil
}
