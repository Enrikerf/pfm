package Entity

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
)

const (
	MaxGas = Pin.MaxDuty
	MinGas = Pin.MinDuty
)

type Engine interface {
	SetGas(gas Pin.Duty)
	SpeedUp()
	MakeLap()
	VelocityControl()
	PositionControl()
	Brake()
	UnBrake()
	Forward()
	Backward()
	GetPosition() int16
	TearDown()
}
type engine struct {
	brakePin   Pin.OutPin
	dirPin     Pin.OutPin
	pwmPin     Pin.PWMPin
	encoder    Encoder
	currentGas Pin.Duty
	forward    bool
}

func NewEngine(
	brakePin Pin.OutPin,
	dirPin Pin.OutPin,
	pwmPin Pin.PWMPin,
	encoder Encoder,
) Engine {
	e := engine{
		brakePin:   brakePin,
		dirPin:     dirPin,
		pwmPin:     pwmPin,
		encoder:    encoder,
		currentGas: 0,
		forward:    true,
	}
	e.watchdog()
	e.initialState()
	return &e
}

func (e *engine) watchdog() {
	go e.encoder.Watchdog()
}

func (e *engine) initialState() {
	e.Forward()
	e.Brake()
	e.encoder.ResetPosition()
}

func (e *engine) MakeLap() {
	e.UnBrake()
	e.SetGas(MinGas)

	if e.forward {
		for e.GetPosition() < 360 {
		}
	} else {
		for e.GetPosition() > -360 {
		}
	}
	e.Brake()
}

func (e *engine) SpeedUp() {
	if e.currentGas < Pin.MinDuty {
		e.currentGas = Pin.MinDuty
	} else {
		e.currentGas++
	}
	e.pwmPin.SetPWM(e.currentGas, Pin.MaxFrequency)
}

func (e *engine) SetGas(gas Pin.Duty) {
	e.currentGas = gas
	e.pwmPin.SetPWM(e.currentGas, Pin.MaxFrequency)
}

func (e *engine) VelocityControl() {
	//TODO implement me
	panic("implement me")
}

func (e *engine) PositionControl() {
	//TODO implement me
	panic("implement me")
}

func (e *engine) Brake() {
	e.pwmPin.StopPWM()
	e.brakePin.Up()
}

func (e *engine) UnBrake() {
	e.brakePin.Down()
}

func (e *engine) Forward() {
	e.forward = true
	e.dirPin.Down()
}

func (e *engine) Backward() {
	e.forward = false
	e.dirPin.Up()
}

func (e *engine) GetPosition() int16 {
	return e.encoder.GetPosition()
}

func (e *engine) TearDown() {
	e.dirPin.TearDown()
	e.brakePin.TearDown()
	e.pwmPin.TearDown()
	e.encoder.TearDown()

}
