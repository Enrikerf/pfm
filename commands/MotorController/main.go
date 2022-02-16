package main

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Config"
)

func main() {
	app := Config.NewApp()
	app.Run()
}
