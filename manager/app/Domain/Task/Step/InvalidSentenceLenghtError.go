package Step

type InvalidSentenceLengthError interface {
	Error() string
}

type invalidSentenceLengthError struct {
}

func (e *invalidSentenceLengthError) Error() string {
	return "step.sentence length must be less than 255"
}

func NewInvalidSentenceLengthError() *invalidSentenceLengthError {
	return &invalidSentenceLengthError{}
}
