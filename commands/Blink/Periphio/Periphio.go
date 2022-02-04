package Periphio

import (
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type Periphio struct {
}

func (p Periphio) Blink() {
	host.Init()

	pinPwm := gpioreg.ByName("18")
	pinEncoderA := gpioreg.ByName("25")
	pinEncoderB := gpioreg.ByName("8")
	pinDir := gpioreg.ByName("7")
	pinBrake := gpioreg.ByName("12")
	t := time.NewTicker(500 * time.Millisecond)
	pinPwm.Out(gpio.Low)
	pinEncoderA.Out(gpio.Low)
	pinEncoderB.Out(gpio.Low)
	pinDir.Out(gpio.Low)
	pinBrake.Out(gpio.Low)
	for l := gpio.Low; ; l = !l {
		pinPwm.Out(l)
		pinEncoderA.Out(l)
		pinEncoderB.Out(l)
		pinDir.Out(l)
		pinBrake.Out(l)
		<-t.C
	}

}
