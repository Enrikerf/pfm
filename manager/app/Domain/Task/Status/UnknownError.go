package Status

import "errors"

type UnknownError interface {
	error
}

func NewUnknownError() UnknownError {
	return errors.New("status not valid")
}
