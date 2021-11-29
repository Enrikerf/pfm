package ApiGrcp

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func Serve() {
	fmt.Println("Starting Command Manager...")
	log.SetFlags(log.LstdFlags | log.Lshortfile) // TODO: parametrise with .env prod/debug :if we crash the go code, we ge the file name and line number

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen at port") //TODO: parametrise port
	}
	var serverOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(serverOptions...)
	task.RegisterTaskServiceServer(grpcServer, TaskServer{})

	//reflection to expose the api doc and commands
	reflection.Register(grpcServer)
	//if error := grpcServer.Serve(listener); error != nil {
	//	log.Fatalf("fatal")
	//}

	go func() {
		fmt.Println("Starting TaskServer...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("fatal")
		}
	}()

	// Wait for control C to exit
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	// Bock until a signal is received
	<-channel
	fmt.Println("Stopping the server")
	grpcServer.Stop()
	fmt.Println("closing the listener")
	listener.Close()
	fmt.Println("End of program")
}