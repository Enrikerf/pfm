package EventDispatcherAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/FillBatchEventHandler"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/TaskEventHandler"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	ResultEvent "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Event"
	TaskEvent "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Event"
)

type EventDispatcherAdapter interface {
	Dispatch(event Event.Event)
}

type eventDispatcherAdapter struct {
	taskEventHandler      TaskEventHandler.UseCase
	fillBatchEventHandler FillBatchEventHandler.UseCase
}

func New(
	taskEventHandler TaskEventHandler.UseCase,
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
	case TaskEvent.TaskCreatedEventName:
		go e.taskEventHandler.Handle(event)
	case ResultEvent.BatchCreatedToBeFilledEventName:
		go e.fillBatchEventHandler.Handle(event)
	}
}
