package Result

type DeleteResult interface {
	Delete(uuid string) error
}
