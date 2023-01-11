package Eraser

import "errors"

type AdapterError interface {
	error
}

func NewAdapterError() TaskNotFoundError {
	return errors.New("error on delete adapter")
}
