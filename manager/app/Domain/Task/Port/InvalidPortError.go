package Port

type InvalidPortError interface {
	Error() string
}

type invalidPortError struct {
}

func (e *invalidPortError) Error() string {
	return "invalid host format"
}

func NewInvalidPortError() *invalidPortError {
	return &invalidPortError{}
}
