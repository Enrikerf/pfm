package Event

type BatchCreatedToBeFilled interface {
	GetName() string
	GetEntityId() string
}

type batchCreatedToBeFilled struct {
	name     string
	entityId string
}

func NewBatchCreatedToBeFilled(entityId string) BatchCreatedToBeFilled {
	self := &batchCreatedToBeFilled{"batch.created_to_be_filled", entityId}
	return self
}

func (t *batchCreatedToBeFilled) GetName() string {
	return t.name
}

func (t *batchCreatedToBeFilled) GetEntityId() string {
	return t.entityId
}
