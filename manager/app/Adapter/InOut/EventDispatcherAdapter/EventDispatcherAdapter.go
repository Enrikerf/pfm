package EventDispatcherAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/FillBatchEventHandler"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/TaskEventHandler"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
)

type EventDispatcherAdapter interface {
	Dispatch(event Event.Event)
}

type eventDispatcherAdapter struct {
	taskEventHandler      TaskEventHandler.TaskEventHandler
	fillBatchEventHandler FillBatchEventHandler.UseCase
}

func New(
	taskEventHandler TaskEventHandler.TaskEventHandler,
	fillBatchEventHandler FillBatchEventHandler.UseCase,
) EventDispatcherAdapter {
	self := &eventDispatcherAdapter{
		taskEventHandler,
		fillBatchEventHandler,
	}
	return self
}

func (e *eventDispatcherAdapter) Dispatch(event Event.Event) {
	switch event.GetName() {
	case "task.created":
		go e.taskEventHandler.Handle(event)
	case "batch.create_to_be_filled":
		go e.fillBatchEventHandler.Handle(event)
	}
}
