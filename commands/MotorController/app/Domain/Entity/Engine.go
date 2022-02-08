package Entity

import "github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"

const (
	MaxGas = Pin.MaxDuty
)

type Engine interface {
	Build()
	Gas()
	VelocityControl()
	PositionControl()
	Brake()
	Dir()
	GetPosition()
}
type engine struct {
	brakePin Pin.OutPin
	dirPin   Pin.OutPin
	pwmPin   Pin.PWMPin
}

func NewEngine() Engine {
	e := engine{}
	e.Build()
	return e
}

func (e engine) Build() {
	e.dirPin.SetStatus()
	e.brakePin.SetStatus()
}

func (e engine) Gas() {
	e.pwmPin.SetPWM(MaxGas/2, Pin.MaxFrequency)
}

func (e engine) VelocityControl() {
	//TODO implement me
	panic("implement me")
}

func (e engine) PositionControl() {
	//TODO implement me
	panic("implement me")
}

func (e engine) Brake() {
	//TODO implement me
	panic("implement me")
}

func (e engine) Dir() {
	//TODO implement me
	panic("implement me")
}

func (e engine) GetPosition() {
	//TODO implement me
	panic("implement me")
}
