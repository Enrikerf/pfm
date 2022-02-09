package Model

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

type pwmPin struct {
	id     string
	status bool
}

func NewPwmPin(id string) Pin.PWMPin {
	_, err := host.Init()
	if err != nil {
		panic("Periph.io Adapter error")
	}
	return &pwmPin{id: id, status: false}
}
func (pwmPin *pwmPin) SetPWM(duty Pin.Duty, frequency physic.Frequency) {

}
func (pwmPin *pwmPin) StopPWM() {

}
