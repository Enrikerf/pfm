package Error

import "errors"

type ManualTaskOnlyCanHave2StepsError interface {
	error
}

func NewManualTaskOnlyCanHave2StepsError() ManualTaskOnlyCanHave2StepsError {
	return errors.New("ManualTaskOnlyCanHave2Steps")
}
