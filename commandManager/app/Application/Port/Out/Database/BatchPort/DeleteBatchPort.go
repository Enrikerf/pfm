package BatchPort

type Delete interface {
	Delete(uuid string) error
}
