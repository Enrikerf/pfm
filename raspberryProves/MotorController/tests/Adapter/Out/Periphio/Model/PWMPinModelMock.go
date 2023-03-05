package ModelTest

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
	"periph.io/x/conn/v3/physic"
)

const (
	minDuty = 1258292 // PWM don't move the EMG30 below this duty
)

type pwmPinMock struct {
	maxDuty int32
	minDuty int32
}

func NewPwmPinMock(maxDuty, minDuty int32) Pin.PWMPin {
	return &pwmPinMock{maxDuty, minDuty}
}

func (outPin *pwmPinMock) TearDown() {
}

func (outPin *pwmPinMock) GetMaxDuty() Pin.Duty {
	return Pin.Duty(outPin.maxDuty)
}

func (outPin *pwmPinMock) GetMinDuty() Pin.Duty {
	return Pin.Duty(outPin.minDuty)
}

func (outPin *pwmPinMock) GetMaxFrequency() Pin.Frequency {
	return Pin.Frequency(10 * physic.KiloHertz)
}

func (pwmPin *pwmPinMock) SetPWM(duty Pin.Duty, frequency Pin.Frequency) {

}

func (pwmPin *pwmPinMock) StopPWM() {

}
