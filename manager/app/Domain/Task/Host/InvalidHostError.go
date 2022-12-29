package Host

type InvalidHostError interface {
	Error() string
}

type invalidHostError struct {
}

func (e *invalidHostError) Error() string {
	return "invalid host format"
}

func NewInvalidHostError() *invalidHostError {
	return &invalidHostError{}
}
