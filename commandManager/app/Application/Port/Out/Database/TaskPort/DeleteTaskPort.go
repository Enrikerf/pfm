package TaskPort

type Delete interface {
	Delete(uuid string) error
}
