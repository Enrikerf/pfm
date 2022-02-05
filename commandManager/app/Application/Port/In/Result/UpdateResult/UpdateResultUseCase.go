package UpdateResult

type UseCase interface {
	Update(command Command) error
}
