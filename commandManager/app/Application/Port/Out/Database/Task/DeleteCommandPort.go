package Task

type DeleteCommand interface {
	Delete(uuid string) error
}
