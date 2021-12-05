package main

import (
	"github.com/Enrikerf/pfm/commandManager/app/Config"
)

func main() {
	var app = Config.GinServer{}
	app.Run()
}
