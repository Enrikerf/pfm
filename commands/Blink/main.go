package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	ledPwm := gpio.NewLedDriver(r, "12")
	ledEncoderA := gpio.NewLedDriver(r, "22")
	ledEncoderB := gpio.NewLedDriver(r, "24")
	ledDir := gpio.NewLedDriver(r, "26")
	ledBrake := gpio.NewLedDriver(r, "28")

	work := func() {
		ledPwm.Toggle()
		ledEncoderA.Toggle()
		ledEncoderB.Toggle()
		ledDir.Toggle()
		ledBrake.Toggle()

	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{ledPwm, ledEncoderA, ledEncoderB, ledDir, ledBrake},
		work,
	)

	robot.Start()
}
