package Pin

import "periph.io/x/conn/v3/physic"

type Duty int32

const (
	MaxFrequency      = 100 * physic.KiloHertz
	MaxDuty      Duty = 1 << 24
)

type PWMPin interface {
	SetPWM(duty Duty, frequency physic.Frequency)
	StopPWM()
}
