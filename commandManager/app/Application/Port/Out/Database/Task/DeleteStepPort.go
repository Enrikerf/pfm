package Task

type DeleteStep interface {
	Delete(uuid string) error
}
