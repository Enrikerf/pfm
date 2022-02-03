package DeleteTask

type UseCase interface {
	Delete(command Command) error
}
