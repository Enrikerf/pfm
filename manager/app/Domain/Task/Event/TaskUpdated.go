package Event

type TaskUpdated interface {
	GetName() string
	GetEntityId() string
}

type taskUpdated struct {
	name     string
	entityId string
}

func NewTaskUpdated(entityId string) TaskUpdated {
	self := &taskCreated{"task.updated", entityId}
	return self
}

func (t *taskUpdated) GetName() string {
	return t.name
}

func (t *taskUpdated) GetEntityId() string {
	return t.entityId
}
