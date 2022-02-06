package DeleteStep

type UseCase interface {
	Delete(command Command) error
}
