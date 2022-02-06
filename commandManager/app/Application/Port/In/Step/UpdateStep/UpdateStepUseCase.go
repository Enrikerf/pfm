package UpdateStep

type UseCase interface {
	Update(command Command) error
}
