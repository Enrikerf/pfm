package MockBatchPort

type DeleteMock struct {
	Error error
}

func (mock DeleteMock) Delete(uuid string) error {
	return mock.Error
}
