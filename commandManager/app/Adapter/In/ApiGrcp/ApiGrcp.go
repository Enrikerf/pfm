package ApiGrcp

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/Controller"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
)

type ApiGrpc struct {
	DB         *gorm.DB
	ServerHost string
	ServerPort string
	GrpcServer *grpc.Server
	Listener   net.Listener
}

func (api *ApiGrpc) Initialize(db *gorm.DB, host string, port string) {
	fmt.Println("Starting Command Manager...")
	api.DB = db
	api.ServerHost = host
	api.ServerPort = port
	api.loadServer()
	api.configControllers()
	api.loadListener()
}

func (api *ApiGrpc) Run() {
	if os.Getenv("APP_DEBUG") == "true" {
		reflection.Register(api.GrpcServer)
	}
	go func() {
		fmt.Println("Starting at: " + api.ServerHost + api.ServerPort)
		if err := api.GrpcServer.Serve(api.Listener); err != nil {
			log.Fatalf("fatal")
		}
	}()
	// Wait for control C to exit
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	// Bock until a signal is received
	<-channel
	fmt.Println("Stopping the server")
	api.GrpcServer.Stop()
	fmt.Println("closing the Listener")
	err := api.Listener.Close()
	if err != nil {
		return
	}
	fmt.Println("End of program")
}

func (api *ApiGrpc) configControllers() {
	var taskAdapter = Task.Adapter{Orm: api.DB}
	var taskService = CreateTask.Service{SavePort: taskAdapter}
	var taskController = Controller.TaskController{SaveTaskUseCase: taskService}
	task.RegisterTaskServiceServer(api.GrpcServer, taskController)
}

func (api *ApiGrpc) loadServer() {
	var serverOptions []grpc.ServerOption
	api.GrpcServer = grpc.NewServer(serverOptions...)
	if os.Getenv("APP_DEBUG") == "true" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

func (api *ApiGrpc) loadListener() {
	listener, err := net.Listen("tcp", api.ServerHost+api.ServerPort)
	if err != nil {
		log.Fatalf("failed to listen at: " + api.ServerHost + api.ServerPort)
	}
	api.Listener = listener
}
