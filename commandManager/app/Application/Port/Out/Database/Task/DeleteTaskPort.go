package Task

type DeleteTask interface {
	Delete(uuid string) error
}
