package FillBatchEventHandler

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Grcp/CallPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type UseCase interface {
	Handle(event Event.Event)
}

func New(
	callBidiPort CallPort.Bidirectional,
	findBatchRepository ResultRepository.FindBatch,
	findTaskRepository TaskRepository.Find,
	saveResultRepository ResultRepository.Save,
) UseCase {
	return &useCase{
		callBidiPort:        callBidiPort,
		findBatchRepository: findBatchRepository,
		finder:              Finder.Finder{FindRepository: findTaskRepository},
		saveResultPort:      saveResultRepository,
	}
}

type useCase struct {
	callBidiPort        CallPort.Bidirectional
	findBatchRepository ResultRepository.FindBatch
	finder              Finder.Finder
	saveResultPort      ResultRepository.Save
}

func (useCase *useCase) Handle(event Event.Event) {
	batchId, err := Result.LoadBatchIdFromString(event.GetEntityId())
	if err != nil {
		return
	}
	batch, err := useCase.findBatchRepository.Find(batchId)
	if err != nil {
		return
	}
	task, err := useCase.finder.Find(batch.GetTaskId())
	if err != nil {
		return
	}
	if len(task.GetSteps()) < 0 {
		return
	}
	err = useCase.callBidiPort.Setup(task.GetHost().GetValue(), task.GetPort().GetValue())
	if err != nil {
		return
	}
	if !useCase.write(task.GetSteps()[0]) {
		useCase.callBidiPort.Close()
		return
	}
	for {
		task, _ := useCase.finder.Find(batch.GetTaskId())
		if task.GetStatus().Value() != Status.Running {
			break
		}
		resultsContent, err := useCase.callBidiPort.Read()
		if err != nil {
			content, err := Content.NewContent(err.Error())
			if err != nil {
				break
			}
			result := Result.New(batchId, content)
			useCase.saveResultPort.Persist(result)
			break
		}
		if len(resultsContent.GetValue()) > 0 {
			result := Result.New(batchId, resultsContent)
			useCase.saveResultPort.Persist(result)
		}
	}
	useCase.callBidiPort.Close()
}

func (useCase *useCase) write(step Step.Step) bool {
	err := useCase.callBidiPort.Write(step)
	if err != nil {
		return false
	}
	return true
}
