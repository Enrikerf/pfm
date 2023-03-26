package Event

type TaskCreated interface {
	GetName() string
	GetEntityId() string
}

type taskCreated struct {
	name     string
	entityId string
}

func NewTaskCreated(entityId string) TaskCreated {
	self := &taskCreated{"task.created", entityId}
	return self
}

func (t *taskCreated) GetName() string {
	return t.name
}

func (t *taskCreated) GetEntityId() string {
	return t.entityId
}
