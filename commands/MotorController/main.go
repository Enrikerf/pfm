package main

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Adapter/Out/Periphio/Model"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"
)

func main() {
	//TODO: this must be in config
	pwmPin := Model.NewPwmPin("18")
	brakePin := Model.NewOutPin("12")
	dirPin := Model.NewOutPin("7")

 
	e:= Entity.NewEngine(brakePin, dirPin, pwmPin)
	e.TearDown()

}
