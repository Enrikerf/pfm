package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Enrikerf/pfm/commands/MotorController/app/Adapter/Out/Periphio/Model"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity"
)

func main() {
	//TODO: this must be in config
	pwmPin := Model.NewPwmPin("18")
	brakePin := Model.NewOutPin("12")
	dirPin := Model.NewOutPin("7")
	encoderPinA := Model.NewEncoderPin("25")
	encoderPinB := Model.NewEncoderPin("8")
	encoder := Model.NewEncoder(encoderPinA, encoderPinB)
	control := Entity.NewControlAlgorithm()
	e := Entity.NewEngine(
		360,
		control,
		brakePin,
		dirPin,
		pwmPin,
		encoder,
	)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	go exitWatchdog(channel, e)

	// e.MakeLap()
	rpm:= 60.0 * 2
	e.RpmControl(rpm)
	e.TearDown()
}

func SetGas(e Entity.Engine) {
	for {
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			panic("io error")
		}
		e.SetGas(Entity.GasLevel(i))
	}
}

func exitWatchdog(channel chan os.Signal, e Entity.Engine) {
	for range channel {
		fmt.Println("tearing down engine")
		e.TearDown()
		fmt.Println("exit succeed")
		os.Exit(1)
	}

}
