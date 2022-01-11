package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Enrikerf/pfm/commandExecutor/app/Adapter/In/ApiGrcp/gen/call"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	call.UnimplementedCallServiceServer
}

func (s Server) CallUnary(ctx context.Context, request *call.CallRequest) (*call.CallResponse, error) {
	println("unary")
	return &call.CallResponse{Result: "result"}, nil
}

func (s Server) CallServerStream(request *call.CallRequest, server call.CallService_CallServerStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) CallClientStream(server call.CallService_CallClientStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) CallBidirectional(server call.CallService_CallBidirectionalServer) error {
	//TODO implement me
	panic("implement me")
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090/echo", "gRPC server endpoint")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Hello world")

	listener, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Fatalf("failed to listen")
	}

	var serverOptions []grpc.ServerOption
	server := grpc.NewServer(serverOptions...)
	call.RegisterCallServiceServer(server, &Server{})

	//reflection to expose the api doc and commands
	reflection.Register(server)
	if error := server.Serve(listener); error != nil {
		log.Fatalf("fatal")
	}

	go func() {
		fmt.Println("Starting Server...")
		if err := server.Serve(listener); err != nil {
			log.Fatalf("fatal")
		}
	}()

	// Wait for control C to exit
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	// Bock until a signal is received
	<-channel
	fmt.Println("Stopping the server")
	server.Stop()
	fmt.Println("closing the listener")
	listener.Close()
	fmt.Println("End of program")
}
