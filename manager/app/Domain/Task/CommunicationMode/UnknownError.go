package CommunicationMode

type UnknownError interface {
	Error() string
}

type unknownError struct {
}

func (e *unknownError) Error() string {
	return "communication mode not valid"
}

func NewUnknownError() *unknownError {
	return &unknownError{}
}
