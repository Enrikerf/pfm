package ExecutionMode

type UnknownError interface {
	Error() string
}

type unknownError struct {
}

func (e *unknownError) Error() string {
	return "execution mode not valid"
}

func NewUnknownError() *unknownError {
	return &unknownError{}
}
