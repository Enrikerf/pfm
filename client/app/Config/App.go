package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/Enrikerf/pfm/commandExecutor/app/Adapter/In/ApiGrcp"
	"github.com/joho/godotenv"
)

type App struct {
	apiGrpc ApiGrcp.ApiGrpc
}

func (server *App) Run() {
	server.loadDotEnv()
	server.loadApiGrpc()

}

func (server *App) loadDotEnv() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}

func (server *App) loadApiGrpc() {
	server.apiGrpc = ApiGrcp.ApiGrpc{}
	server.apiGrpc.Initialize(
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	go server.apiGrpc.Run()

}
