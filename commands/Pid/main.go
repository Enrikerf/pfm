package main

import (
	"fmt"
	"os"
	"os/signal"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

func main() {
	_, err := host.Init()
	if err != nil {
		return
	}

	pinPwm := gpioreg.ByName("18")
	pinEncoderA := gpioreg.ByName("25")
	pinEncoderB := gpioreg.ByName("8")
	pinDir := gpioreg.ByName("7")
	pinBrake := gpioreg.ByName("12")

	pinEncoderA.In(gpio.PullUp, gpio.BothEdges)
	pinEncoderB.In(gpio.PullUp, gpio.BothEdges)
	pinDir.Out(gpio.High)
	pinBrake.Out(gpio.Low)
	if err := pinPwm.PWM(gpio.DutyMax/3, 100*physic.KiloHertz); err != nil {
		fmt.Println(err)
	}
	go encoder(pinEncoderA, pinEncoderB)
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	fmt.Println("prev to exit")
	<-channel
	fmt.Println("exiting")
	pinEncoderA.Out(gpio.Low)
	pinEncoderB.Out(gpio.Low)
	pinDir.Out(gpio.Low)
	pinBrake.Out(gpio.Low)
	pinPwm.PWM(0, 0)
	if err := pinPwm.Halt(); err != nil {
		fmt.Println(err)
	}
	pinPwm.Out(gpio.Low)
	os.Exit(1)
	fmt.Println()
	fmt.Println("exit")
}

func encoder(pinEncoderA gpio.PinIO, pinEncoderB gpio.PinIO) {
	encoderPinALast := gpio.Low
	encoderPos := 0
	for {
		// pinEncoderA.WaitForEdge(-1)
		// fmt.Printf("-> %s\n", pinEncoderA.Read())
		n := pinEncoderA.Read()
		if encoderPinALast == gpio.Low && n == gpio.High {
			if pinEncoderB.Read() == gpio.Low {
				encoderPos--
			} else {
				encoderPos++
			}
			fmt.Println(encoderPos)
		}
		encoderPinALast = n
	}
}

// var halt = make(chan os.Signal)
// 	signal.Notify(halt, syscall.SIGTERM)
// 	signal.Notify(halt, syscall.SIGINT)

// 	go func() {
// 		select {
// 		case <-halt:
// 			fmt.Println("exit")
// 			pinPwm.Out(gpio.Low)
// 			pinEncoderA.Out(gpio.Low)
// 			pinEncoderB.Out(gpio.Low)
// 			pinDir.Out(gpio.Low)
// 			pinBrake.Out(gpio.Low)
// 			if err := pinPwm.Halt(); err != nil {
// 				fmt.Println(err)
// 			}
// 			os.Exit(1)
// 		}
// 	}()
