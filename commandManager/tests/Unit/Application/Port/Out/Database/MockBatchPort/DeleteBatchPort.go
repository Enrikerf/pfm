package MockBatchPort

type DeleteMock struct {
	Err error
}

func (mock DeleteMock) Delete(uuid string) error {
	return mock.Err
}
