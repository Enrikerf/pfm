package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Enrikerf/pfm/commands/MotorController/app/Config"
)

func main() {

	app := Config.NewApp()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	go exitWatchdog(channel, app)
	app.Run()

}

func exitWatchdog(channel chan os.Signal, app Config.App) {
	for range channel {
		var i string
		fmt.Println("\nDo you want to exit? (y/n)")
		_, err := fmt.Scanf("%s", &i)
		if err != nil {
			fmt.Println(err)
			app.TearDown()
			panic(err)
		}
		if i == "y" {
			fmt.Println("tearing down engine")
			app.TearDown()
			fmt.Println("exit succeed")
			os.Exit(1)
		} else {

		}
	}
}
