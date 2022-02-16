package Config

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Adapter/In/Console"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Adapter/Out/Periphio/Model"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Application/Port/In/ManageEngine"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"
)

type App interface {
	Run()
}
type app struct {
	console Console.Console
}

func NewApp() App {
	pwmPin := Model.NewPwmPin("18")
	brakePin := Model.NewOutPin("12")
	dirPin := Model.NewOutPin("7")
	encoderPinA := Model.NewEncoderPin("25")
	encoderPinB := Model.NewEncoderPin("8")
	encoder := Model.NewEncoder(encoderPinA, encoderPinB)
	control := Entity.NewControlAlgorithm()
	engine := Entity.NewEngine(
		360,
		control,
		brakePin,
		dirPin,
		pwmPin,
		encoder,
	)
	service := ManageEngine.Service{Engine: engine}
	console := Console.NewConsole(&service)
	return &app{console: console}
}

func (app *app) Run() {
	app.console.Run()
}
