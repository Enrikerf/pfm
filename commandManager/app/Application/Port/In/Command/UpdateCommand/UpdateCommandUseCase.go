package UpdateCommand

type UseCase interface {
	Update(command Command) error
}
