package UpdateBatch

type UseCase interface {
	Update(command Command) error
}
