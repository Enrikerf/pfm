package Config

import (
	"fmt"
	"os"

	"github.com/Enrikerf/pfm/commands/MotorController/app/Adapter/Out/Periphio/Model"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"
)

type App interface {
	Run()
	TearDown()
}
type app struct {
	engine Entity.Engine
	isCommandRunning chan bool
}

func NewApp() App {
	pwmPin := Model.NewPwmPin("18")
	brakePin := Model.NewOutPin("12")
	dirPin := Model.NewOutPin("7")
	encoderPinA := Model.NewEncoderPin("25")
	encoderPinB := Model.NewEncoderPin("8")
	encoder := Model.NewEncoder(encoderPinA, encoderPinB)
	control := Entity.NewControlAlgorithm()
	return &app{engine: Entity.NewEngine(
		360,
		control,
		brakePin,
		dirPin,
		pwmPin,
		encoder,
	)}
}

func (app *app) Run() {

	var i int
	for {
		fmt.Println("commands:")
		fmt.Println("\t1) MakeLap")
		fmt.Println("\t2) up rpm control")
		fmt.Println("\t3) down rpm control")
		fmt.Println("\t4) set gas")
		fmt.Println("\t5) step response")
		fmt.Printf("choose: ")
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println(err)
			app.TearDown()
			panic(err)
		}
		switch i {
		case 1:
			app.engine.MakeLap()
		case 2:
			rpm := 60.0 * 2
			go app.engine.RpmControl(rpm)
		case 3:
			
		case 4:
			app.setGas()
		case 5:
			app.engine.StepResponse()
		case -1:
			app.engine.TearDown()
			os.Exit(1)
			break
		default:
			fmt.Println("- command not found")
		}
	}
}

func (app *app) TearDown() {
	app.engine.TearDown()
}

func (app *app) setGas() {
	app.engine.UnBrake()
	for {
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println(err)
			app.TearDown()
			panic(err)
		}
		app.engine.SetGas(Entity.GasLevel(i))
	}
}
