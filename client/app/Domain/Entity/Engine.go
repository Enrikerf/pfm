package Entity

import (
	"fmt"
	"math"
	"time"

	"github.com/Enrikerf/pfm/commandExecutor/app/Domain/Entity/Pin"
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
	GetCurrentAngularSpeed() float64
	TearDown()
	InitialState()
}

type engine struct {
	encoderSlots        int
	resolution          float64
	controlAlgorithm    ControlAlgorithm
	brakePin            Pin.OutPin
	dirPin              Pin.OutPin
	pwmPin              Pin.PWMPin
	encoder             Encoder
	currentGas          GasLevel
	forward             bool
	isControlRunning    chan bool
	currentAngularSpeed float64
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
		encoderSlots:        encoderSlots,
		resolution:          360.0 / float64(encoderSlots),
		controlAlgorithm:    controlAlgorithm,
		brakePin:            brakePin,
		dirPin:              dirPin,
		pwmPin:              pwmPin,
		encoder:             encoder,
		currentGas:          0,
		forward:             true,
		isControlRunning:    make(chan bool, 1),
		currentAngularSpeed: 0,
	}
	e.watchdog()
	e.InitialState()
	return &e
}

func (e *engine) watchdog() {
	go e.encoder.Watchdog()
}

func (e *engine) InitialState() {
	e.Forward()
	e.Brake()
	e.encoder.ResetPosition()
}

func (e *engine) MakeLap() {
	e.InitialState()
	e.UnBrake()
	e.SetGas(GasLevel(e.pwmPin.GetMinDuty()))

	if e.forward {
		for e.GetPosition() < 360 {
			//fmt.Println(e.GetPosition())
		}
	} else {
		for e.GetPosition() > -360 {
			//fmt.Println(e.GetPosition())
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
		e.currentAngularSpeed = radianPerSecod
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
	radPerSecondGoal := 60 * (2 * math.Pi / 60)
	e.controlAlgorithm.SetGoal(radPerSecondGoal)
	e.controlAlgorithm.SetP(2000000)
	e.controlAlgorithm.SetI(100)
	e.controlAlgorithm.SetD(0)
	e.controlAlgorithm.SetSampleTime(sampleTime.Seconds())

	prevAngle := 0.0
	e.brakePin.Down()
	ticker := time.NewTicker(sampleTime)
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
				fmt.Println("prevAngle", prevAngle)
				angle := e.resolution * math.Abs(float64(e.encoder.GetPosition()))
				fmt.Println("angle:", angle)
				fmt.Println("sampleTime.Seconds():", sampleTime.Seconds())
				degreesPerSecod := (angle - prevAngle) / sampleTime.Seconds()
				fmt.Println("degreesPerSecod:", degreesPerSecod)
				radianPerSecod := degreesPerSecod * math.Pi / 180
				fmt.Println("radianPerSecod:", radianPerSecod)
				e.currentAngularSpeed = radianPerSecod
				pidOrig := e.controlAlgorithm.Calculate(radianPerSecod)
				pidReescalated := e.reescalePid(pidOrig)
				fmt.Println("-------------------------------")
				e.pwmPin.SetPWM(Pin.Duty(pidReescalated), e.pwmPin.GetMaxFrequency())
				prevAngle = angle
				if len(e.isControlRunning) == 0 {
					e.pwmPin.SetPWM(Pin.Duty(0), e.pwmPin.GetMaxFrequency())
					e.brakePin.Up()
					ticker.Stop()
					done <- true
					fmt.Println("Ticker stopped")
				}
            }
        }
    }()



	// e.brakePin.Down()
	// for range time.Tick(sampleTime) {
	// 	fmt.Println("prevAngle", prevAngle)
	// 	angle := e.resolution * math.Abs(float64(e.encoder.GetPosition()))
	// 	fmt.Println("angle:", angle)
	// 	fmt.Println("sampleTime.Seconds():", sampleTime.Seconds())
	// 	degreesPerSecod := (angle - prevAngle) / sampleTime.Seconds()
	// 	fmt.Println("degreesPerSecod:", degreesPerSecod)
	// 	radianPerSecod := degreesPerSecod * math.Pi / 180
	// 	fmt.Println("radianPerSecod:", radianPerSecod)
	// 	e.currentAngularSpeed = radianPerSecod
	// 	pidOrig := e.controlAlgorithm.Calculate(radianPerSecod)
	// 	pidReescalated := e.reescalePid(pidOrig)
	// 	fmt.Println("-------------------------------")
	// 	e.pwmPin.SetPWM(Pin.Duty(math.Abs(pidReescalated)), e.pwmPin.GetMaxFrequency())
	// 	// if pidReescalated <0{
	// 	// 	e.Backward()
	// 	// }else{
	// 	// 	e.Forward()
	// 	// }
	// 	prevAngle = angle
	// 	if len(e.isControlRunning) == 0 {
	// 		e.pwmPin.SetPWM(Pin.Duty(0), e.pwmPin.GetMaxFrequency())
	// 		e.brakePin.Up()
	// 		break
	// 	}
	// 	// time.Sleep(time.Millisecond * 100)
	// }
}

func (e *engine) reescalePid(pid float64) float64 {
	fmt.Println("\tsubfuc:")
	fmt.Println("\t\tpid:", pid)
	reescalePid := pid
	// minV := 17 * 2 * math.Pi / 60 //Probado con test que eso es a lo minimo que va
	if math.Abs(pid) > float64(e.pwmPin.GetMaxDuty()){
		if pid < 0{
			fmt.Println("\t\t changed to min")
			reescalePid = -float64(e.pwmPin.GetMaxDuty())
		}else{
			fmt.Println("\t\t changed to max")
			reescalePid = float64(e.pwmPin.GetMaxDuty())
		}
	}
	rangeMin:= -float64(e.pwmPin.GetMaxDuty())
	rangeMax:=  float64(e.pwmPin.GetMaxDuty())
	targetMin:= 0.0
	targetMax:= float64(e.pwmPin.GetMaxDuty())
	reescalePid = ((reescalePid-rangeMin)/(rangeMax-rangeMin))*(targetMax-targetMin) + targetMin
	fmt.Println("\t\treescalePid:", reescalePid)
	return reescalePid
}

func (e *engine) StopRmpControl() {
	if len(e.isControlRunning) > 0 {
		<-e.isControlRunning
	} else {
		fmt.Println("no command running")
	}
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

func (e *engine) GetCurrentAngularSpeed() float64 {
	return e.currentAngularSpeed
}
