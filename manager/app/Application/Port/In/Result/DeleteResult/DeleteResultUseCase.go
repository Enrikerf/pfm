package DeleteResult

type UseCase interface {
	Delete(command Command) error
}
