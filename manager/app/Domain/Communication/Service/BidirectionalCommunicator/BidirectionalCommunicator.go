package BidirectionalCommunicator

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type BidirectionalCommunicator interface {
	Communicate(task Task.Task, batch Result.Batch)
}

func New(
	find TaskRepository.Find,
	callBidiPort Repository.Bidirectional,
	saveResultRepository ResultRepository.Save,
) BidirectionalCommunicator {
	return &bidirectionalCommunicator{
		finder:         Finder.Finder{FindRepository: find},
		callBidiPort:   callBidiPort,
		saveResultPort: saveResultRepository,
	}
}

type bidirectionalCommunicator struct {
	finder         Finder.Finder
	callBidiPort   Repository.Bidirectional
	saveResultPort ResultRepository.Save
}

func (b *bidirectionalCommunicator) Communicate(task Task.Task, batch Result.Batch) {

	err := b.callBidiPort.Setup(task.GetHost().GetValue(), task.GetPort().GetValue())
	if err != nil {
		return
	}
	if !b.write(task.GetSteps()[0]) {
		b.callBidiPort.Close()
		return
	}
	for {
		task, _ := b.finder.Find(batch.GetTaskId())
		if task.GetStatus().Value() != Status.Running {
			break
		}
		resultsContent, err := b.callBidiPort.Read()
		if err != nil {
			content, err := Content.NewContent(err.Error())
			if err != nil {
				break
			}
			result := Result.New(batch.GetId(), content)
			b.saveResultPort.Persist(result)
			break
		}
		if len(resultsContent.GetValue()) > 0 {
			result := Result.New(batch.GetId(), resultsContent)
			b.saveResultPort.Persist(result)
		}
	}
	b.callBidiPort.Close()
}

func (b *bidirectionalCommunicator) write(step Step.Step) bool {
	err := b.callBidiPort.Write(step)
	if err != nil {
		return false
	}
	return true
}
