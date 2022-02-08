package Pin

const (
	HighStatus = true
	LowStatus  = false
)

type OutPin interface {
	GetId() string
	GetStatus() bool
	SetStatus()
}

type outPin struct {
	id     string
	status bool
}

func NewPin(id string) OutPin {
	return &outPin{id: id, status: false}
}

func (outPin outPin) SetStatus() {
	//TODO implement me
	panic("implement me")
}

func FullDefinedPin(id string, status bool) *outPin {
	return &outPin{id: id, status: status}
}

func (outPin *outPin) GetId() string {
	return outPin.id
}

func (outPin *outPin) GetStatus() bool {
	return outPin.status
}
