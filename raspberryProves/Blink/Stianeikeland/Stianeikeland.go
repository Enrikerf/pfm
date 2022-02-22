package Stianeikeland

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type Stianeikeland struct {
}

func (s Stianeikeland) Blink() {
	pinPwm := rpio.Pin(18)
	pinEncoderA := rpio.Pin(25)
	pinEncoderB := rpio.Pin(8)
	pinDir := rpio.Pin(7)
	pinBrake := rpio.Pin(12)
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pinPwm.Output()
	pinEncoderA.Output()
	pinEncoderB.Output()
	pinDir.Output()
	pinBrake.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pinPwm.Toggle()
		pinEncoderA.Toggle()
		pinEncoderB.Toggle()
		pinDir.Toggle()
		pinBrake.Toggle()
		time.Sleep(time.Second)
	}

}
