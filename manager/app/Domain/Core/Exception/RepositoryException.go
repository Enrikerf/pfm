package Exception

import "errors"

type RepositoryException interface {
	error
}

func NewRepositoryException() RepositoryException {
	return errors.New("repository not handled exception")
}
