package DeleteCommand

type UseCase interface {
	Delete(command Command) error
}
