package Error

import "errors"

type UnaryTaskCanOnlyHaveOneStepError interface {
	error
}

func NewUnaryTaskCanOnlyHaveOneStepError() UnaryTaskCanOnlyHaveOneStepError {
	return errors.New("UnaryTaskCanOnlyHaveOneStepError")
}
