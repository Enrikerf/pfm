package MockTaskPort

type MockDeletePort struct {
	Error error
}

func (mock MockDeletePort) Delete(uuid string) error {
	return mock.Error
}
