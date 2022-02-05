package Result

type Delete interface {
	Delete(uuid string) error
}
