package StepPort

type DeleteStep interface {
	Delete(uuid string) error
}
