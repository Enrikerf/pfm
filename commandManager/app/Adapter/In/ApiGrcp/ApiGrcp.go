package ApiGrcp

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/Controller"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
)

func Serve() {

	var err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	dbServerHost := os.Getenv("SERVER_HOST")
	dbServerPort := os.Getenv("SERVER_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbUrl, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	var taskAdapter = Task.Adapter{Orm: db}
	var taskService = CreateTask.Service{SavePort: taskAdapter}
	var taskController = Controller.TaskController{SaveTaskUseCase: taskService}
	fmt.Println("Starting Command Manager...")
	log.SetFlags(log.LstdFlags | log.Lshortfile) // TODO: parametrise with .env prod/debug :if we crash the go code, we ge the file name and line number

	listener, err := net.Listen("tcp", dbServerHost+dbServerPort)
	if err != nil {
		log.Fatalf("failed to listen at: " + dbServerHost + dbServerPort)
	}
	var serverOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(serverOptions...)
	task.RegisterTaskServiceServer(grpcServer, taskController)

	//reflection to expose the api doc and commands
	reflection.Register(grpcServer)
	//if error := grpcServer.Serve(listener); error != nil {
	//	log.Fatalf("fatal")
	//}

	go func() {
		fmt.Println("Starting TaskController...")
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
