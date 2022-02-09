package Model

import (
	"fmt"
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Entity/Pin"
	"periph.io/x/host/v3"
)

type outPin struct {
	id     string
	status bool
}

func NewOutPin(id string) Pin.OutPin {
	_, err := host.Init()
	if err != nil {
		panic("Periph.io Adapter error")
	}
	return &outPin{id: id, status: false}
}

func (outPin *outPin) Up() {
	outPin.status = true
	fmt.Printf("Pin %v: %v\n", outPin.id, outPin.status)
}

func (outPin *outPin) Down() {
	outPin.status = false
	fmt.Printf("Pin %v: %v\n", outPin.id, outPin.status)
}
