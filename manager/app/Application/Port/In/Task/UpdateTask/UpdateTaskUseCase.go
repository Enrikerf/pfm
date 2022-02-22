package UpdateTask

type UseCase interface {
	Update(command Command) error
}
