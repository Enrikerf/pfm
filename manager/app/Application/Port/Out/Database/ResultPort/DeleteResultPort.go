package ResultPort

type Delete interface {
	Delete(uuid string) error
}
