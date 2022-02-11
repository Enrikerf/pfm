package Pin

import "periph.io/x/conn/v3/physic"

type Duty int32

const (
	MaxFrequency      = 100 * physic.KiloHertz
	MaxDuty      Duty = 1 << 24
	MinDuty      Duty = 1258292
)

type PWMPin interface {
	SetPWM(duty Duty, frequency physic.Frequency)
	StopPWM()
	TearDown()
}
