package Entity

import "github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"

const (
	MaxGas = Pin.MaxDuty
)

type Engine interface {
	Gas()
	VelocityControl()
	PositionControl()
	Brake()
	UnBrake()
	Forward()
	Backward()
	GetPosition()
	TearDown()
}
type engine struct {
	brakePin Pin.OutPin
	dirPin   Pin.OutPin
	pwmPin   Pin.PWMPin
}

func NewEngine(
	brakePin Pin.OutPin,
	dirPin Pin.OutPin,
	pwmPin Pin.PWMPin,
) Engine {
	e := engine{
		brakePin: brakePin,
		dirPin:   dirPin,
		pwmPin:   pwmPin,
	}
	e.setup()
	return e
}

func (e engine) setup() {
	e.dirPin.Up()
	e.brakePin.Up()
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
	e.brakePin.Up()
}

func (e engine) UnBrake() {
	e.brakePin.Down()
}

func (e engine) Forward() {
	e.dirPin.Up()
}

func (e engine) Backward() {
	e.dirPin.Down()
}

func (e engine) GetPosition() {
	e.dirPin.Down()
}

func (e engine) TearDown(){
	e.dirPin.TearDown()
	e.brakePin.TearDown()
}
