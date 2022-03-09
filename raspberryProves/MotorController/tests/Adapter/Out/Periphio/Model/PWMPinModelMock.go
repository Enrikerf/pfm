package ModelTest

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/physic"
)

const (
	minDuty = 1258292 // PWM don't move the EMG30 below this duty
)

type pwmPinMock struct {
}

func NewPwmPinMock() Pin.PWMPin {
	return &pwmPinMock{}
}

func (outPin *pwmPinMock) TearDown() {
}

func (outPin *pwmPinMock) GetMaxDuty() Pin.Duty {
	return Pin.Duty(gpio.DutyMax)
}

func (outPin *pwmPinMock) GetMinDuty() Pin.Duty {
	return Pin.Duty(minDuty)
}

func (outPin *pwmPinMock) GetMaxFrequency() Pin.Frequency {
	return Pin.Frequency(10 * physic.KiloHertz)
}

func (pwmPin *pwmPinMock) SetPWM(duty Pin.Duty, frequency Pin.Frequency) {

}

func (pwmPin *pwmPinMock) StopPWM() {

}
