package Entity

type Encoder interface {
	Watchdog()
	ResetPosition()
	GetPosition() int64
	TearDown()
}
