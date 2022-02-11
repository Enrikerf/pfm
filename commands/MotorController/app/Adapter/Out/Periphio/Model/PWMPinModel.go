package Model

import (
	"fmt"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

type pwmPin struct {
	periphPin gpio.PinIO
	id        string
	status    bool
}

func NewPwmPin(id string) Pin.PWMPin {
	_, err := host.Init()
	if err != nil {
		panic("Periph.io Adapter error")
	}
	e := pwmPin{id: id, status: false, periphPin: gpioreg.ByName(id)}
	return &e
}

func (outPin *pwmPin) TearDown() {
	outPin.periphPin.Halt()
	outPin.periphPin.DefaultPull()
	outPin.periphPin.Out(gpio.Low)
}

func (pwmPin *pwmPin) SetPWM(duty Pin.Duty, frequency physic.Frequency) {
	if err := pwmPin.periphPin.PWM(gpio.Duty(duty), 10*physic.KiloHertz); err != nil {
		fmt.Println(err)
	}
}

func (pwmPin *pwmPin) StopPWM() {
	if err := pwmPin.periphPin.PWM(0, 100*physic.KiloHertz); err != nil {
		panic("%d: pwm panic")
	}
}
