package main

import (
	"fmt"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
)

func main() {
	x := Pin.NewPin("18")
	helloPin(x)
	b := Pin.FullDefinedPin("18", Pin.HighStatus)
	helloPin(b)
}

func helloPin(pin Pin.OutPin) {
	fmt.Println(pin.GetId())
}
