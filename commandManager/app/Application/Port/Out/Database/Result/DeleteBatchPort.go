package Result

type DeleteBatch interface {
	Delete(uuid string) error
}
