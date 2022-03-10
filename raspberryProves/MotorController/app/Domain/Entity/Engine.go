package Entity

import (
	"fmt"
	"math"
	"time"

	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
)

type GasLevel int32

type Engine interface {
	SetGas(gas GasLevel)
	StepResponse()
	SpeedUp()
	MakeLap()
	RpmControl(goal float64)
	StopRmpControl()
	PositionControl()
	Brake()
	UnBrake()
	Forward()
	Backward()
	GetPosition() int16
	TearDown()
	ReescalePid(pid float64) int32
}

type engine struct {
	encoderSlots     int
	resolution       float64
	controlAlgorithm ControlAlgorithm
	brakePin         Pin.OutPin
	dirPin           Pin.OutPin
	pwmPin           Pin.PWMPin
	encoder          Encoder
	currentGas       GasLevel
	forward          bool
	isControlRunning chan bool
}

//TODO: make singleton
func NewEngine(
	encoderSlots int,
	controlAlgorithm ControlAlgorithm,
	brakePin Pin.OutPin,
	dirPin Pin.OutPin,
	pwmPin Pin.PWMPin,
	encoder Encoder,
) Engine {
	e := engine{
		encoderSlots:     encoderSlots,
		resolution:       360.0 / float64(encoderSlots),
		controlAlgorithm: controlAlgorithm,
		brakePin:         brakePin,
		dirPin:           dirPin,
		pwmPin:           pwmPin,
		encoder:          encoder,
		currentGas:       0,
		forward:          true,
		isControlRunning: make(chan bool, 1),
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
	e.initialState()
	e.UnBrake()
	e.SetGas(GasLevel(e.pwmPin.GetMinDuty()))

	if e.forward {
		for e.GetPosition() < 360 {
			fmt.Println(e.GetPosition())
		}
	} else {
		for e.GetPosition() > -360 {
			fmt.Println(e.GetPosition())
		}
	}
	e.Brake()
}

func (e *engine) SpeedUp() {
	if e.currentGas < GasLevel(e.pwmPin.GetMinDuty()) {
		e.currentGas = GasLevel(e.pwmPin.GetMinDuty())
	} else {
		e.currentGas++
	}
	e.pwmPin.SetPWM(Pin.Duty(e.currentGas), e.pwmPin.GetMaxFrequency())
}

func (e *engine) SetGas(gas GasLevel) {
	e.currentGas = gas
	e.pwmPin.SetPWM(Pin.Duty(e.currentGas), e.pwmPin.GetMaxFrequency())
}

func (e *engine) StepResponse() {
	sampleTime := time.Millisecond * 10
	prevAngle := 0.0
	e.encoder.ResetPosition()
	e.brakePin.Down()
	cont := 0
	for range time.Tick(sampleTime) {
		angle := e.resolution * math.Abs(float64(e.encoder.GetPosition()))
		degreesPerSecod := (angle - prevAngle) / sampleTime.Seconds()
		radianPerSecod := degreesPerSecod * math.Pi / 180
		fmt.Println(" ", radianPerSecod)
		prevAngle = angle
		if cont == 1 {
			e.pwmPin.SetPWM(e.pwmPin.GetMinDuty()*5, e.pwmPin.GetMaxFrequency())
		}
		cont++
	}
}

func (e *engine) RpmControl(goal float64) {
	go e.contrlLoop(goal)
}

func (e *engine) contrlLoop(goal float64) {
	e.isControlRunning <- true
	sampleTime := time.Millisecond * 10
	radPerSecondGoal := goal * (2 * math.Pi / 60)
	e.controlAlgorithm.SetGoal(radPerSecondGoal)
	e.controlAlgorithm.SetP(1)
	e.controlAlgorithm.SetI(10)
	e.controlAlgorithm.SetD(0)
	e.controlAlgorithm.SetLowerConstraint(float64(e.pwmPin.GetMinDuty()))
	e.controlAlgorithm.SetUpperConstraint(float64(e.pwmPin.GetMaxDuty()))
	e.controlAlgorithm.SetSampleTime(sampleTime.Seconds())

	prevAngle := 0.0

	e.brakePin.Down()
	for range time.Tick(sampleTime) {
		angle := e.resolution * math.Abs(float64(e.encoder.GetPosition()))
		degreesPerSecod := (angle - prevAngle) / sampleTime.Seconds()
		radianPerSecod := degreesPerSecod * math.Pi / 180
		pidOrig := e.controlAlgorithm.Calculate(radianPerSecod)
		if pidOrig == -1 { //TODO: control error properly
			e.pwmPin.SetPWM(Pin.Duty(0), e.pwmPin.GetMaxFrequency())
			e.brakePin.Up()
			<-e.isControlRunning
			break
		}
		e.pwmPin.SetPWM(Pin.Duty(pidOrig), e.pwmPin.GetMaxFrequency())
		prevAngle = angle
		if len(e.isControlRunning) == 0 {
			e.pwmPin.SetPWM(Pin.Duty(0), e.pwmPin.GetMaxFrequency())
			e.brakePin.Up()
			break
		}
	}
}

func (e *engine) StopRmpControl() {
	if len(e.isControlRunning) > 0 {
		<-e.isControlRunning
	} else {
		fmt.Println("no command running")
	}
}

func (e *engine) ReescalePid(pid float64) int32 {
	if pid < float64(e.pwmPin.GetMinDuty()) {
		return int32(e.pwmPin.GetMinDuty())
	}
	if pid > float64(e.pwmPin.GetMaxDuty()) {
		return int32(e.pwmPin.GetMaxDuty())
	}
	return int32(pid)
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
