package ApiGrcp

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/Controller"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ShowTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/UpdateTask"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type ApiGrpc struct {
	createTaskUseCase CreateTask.UseCase
	listTasksUseCase  ListTasks.UseCase
	showTaskUseCase   ShowTask.UseCase
	deleteTaskUseCase DeleteTask.UseCase
	updateTaskUseCase UpdateTask.UseCase
	serverHost        string
	serverPort        string
	grpcServer        *grpc.Server
	listener          net.Listener
}

func (api *ApiGrpc) Initialize(
	createTaskUseCase CreateTask.UseCase,
	listTasksUseCase ListTasks.UseCase,
	showTaskUseCase ShowTask.UseCase,
	deleteTaskUseCase DeleteTask.UseCase,
	updateTaskUseCase UpdateTask.UseCase,
	host string,
	port string,
) {
	fmt.Println("Starting Name Manager...")
	api.createTaskUseCase = createTaskUseCase
	api.listTasksUseCase = listTasksUseCase
	api.showTaskUseCase = showTaskUseCase
	api.deleteTaskUseCase = deleteTaskUseCase
	api.updateTaskUseCase = updateTaskUseCase
	api.serverHost = host
	api.serverPort = port
	api.loadServer()
	api.configControllers()
	api.loadListener()
}

func (api *ApiGrpc) Run() {
	if os.Getenv("APP_DEBUG") == "true" {
		reflection.Register(api.grpcServer)
	}
	go func() {
		fmt.Println("Starting at: " + api.serverHost + api.serverPort)
		if err := api.grpcServer.Serve(api.listener); err != nil {
			log.Fatalf("fatal")
		}
	}()
	// Wait for control C to exit
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	// Bock until a signal is received
	<-channel
	fmt.Println("Stopping the server")
	api.grpcServer.Stop()
	fmt.Println("closing the Listener")
	err := api.listener.Close()
	if err != nil {
		return
	}
	fmt.Println("End of program")
}

func (api *ApiGrpc) configControllers() {
	var taskController = Controller.TaskController{
		SaveTaskUseCase:   api.createTaskUseCase,
		ListTasksUseCase:  api.listTasksUseCase,
		ShowTaskUseCase:   api.showTaskUseCase,
		DeleteTaskUseCase: api.deleteTaskUseCase,
		UpdateTaskUseCase: api.updateTaskUseCase,
	}
	task.RegisterTaskServiceServer(api.grpcServer, taskController)
}

func (api *ApiGrpc) loadServer() {
	var serverOptions []grpc.ServerOption
	api.grpcServer = grpc.NewServer(serverOptions...)
	if os.Getenv("APP_DEBUG") == "true" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

func (api *ApiGrpc) loadListener() {
	listener, err := net.Listen("tcp", api.serverHost+api.serverPort)
	if err != nil {
		log.Fatalf("failed to listen at: " + api.serverHost + api.serverPort)
	}
	api.listener = listener
}
