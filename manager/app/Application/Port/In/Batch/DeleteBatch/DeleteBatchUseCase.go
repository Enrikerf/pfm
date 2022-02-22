package DeleteBatch

type UseCase interface {
	Delete(command Command) error
}
