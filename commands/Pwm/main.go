package main

import (
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
	"time"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Println(err)
	}

	pwm0 := gpioreg.ByName("18")
	defer pwm0.Halt()
	if err := pwm0.Out(gpio.Low); err != nil {
		log.Fatalf("Failed to parse green led to low: %s", err.Error())
	}

	i, err := gpio.ParseDuty("10%")
	if err != nil {
		log.Fatalf("Failed to parse duty cycle: %s", err.Error())
	}

	step := i / 2

	for {
		log.Printf("Raising PWM to: %v duty", i)
		if err := pwm0.PWM(i, 10*physic.KiloHertz); err != nil {
			log.Printf("Failed to set PWM: %s", err.Error())
		}

		i += step
		if i > gpio.DutyMax {
			break
		}

		time.Sleep(1000 * time.Millisecond)
	}

	var halt = make(chan os.Signal)
	signal.Notify(halt, syscall.SIGTERM)
	signal.Notify(halt, syscall.SIGINT)

	go func() {
		select {
		case <-halt:
			if err := pwm0.Halt(); err != nil {
				fmt.Println(err)
			}
			os.Exit(1)
		}
	}()
}