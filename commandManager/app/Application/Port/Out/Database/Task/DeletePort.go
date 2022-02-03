package Task

type Delete interface {
	Delete(uuid string) error
}
