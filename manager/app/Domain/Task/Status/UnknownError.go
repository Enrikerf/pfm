package Status

type UnknownError interface {
	Error() string
}

type unknownError struct {
}

func (e *unknownError) Error() string {
	return "status not valid"
}

func NewUnknownError() *unknownError {
	return &unknownError{}
}
