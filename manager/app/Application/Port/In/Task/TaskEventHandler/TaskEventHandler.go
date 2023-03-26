package TaskEventHandler

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type TaskEventHandler interface {
	Handle(created Event.TaskCreated)
}

func New(findRepository Repository.Find) TaskEventHandler {
	return &taskEventHandler{
		Finder.Finder{FindRepository: findRepository},
		NewLooper(),
	}
}

type taskEventHandler struct {
	finder Finder.Finder
	looper Looper
}

func (useCase *taskEventHandler) Handle(taskChanged Event.TaskCreated) {
	taskId, err := Task.LoadIdFromString(taskChanged.GetEntityId())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	task, err := useCase.finder.Find(taskId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if task.GetExecutionMode() != ExecutionMode.Automatic && task.GetStatus().Value() != Status.Pending {
		return
	}
	if !useCase.looper.IsEnabled() {
		useCase.looper.Enable()
	}
}
