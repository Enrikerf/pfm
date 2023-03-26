package Error

import "errors"

type RepositoryError interface {
	error
}

func NewRepositoryError() RepositoryError {
	return errors.New("repository not handled exception")
}
