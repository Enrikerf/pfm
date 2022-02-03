package Config

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/LoopManager"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Grcp/Call"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Call/Loop"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ShowTask"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type App struct {
	ApiGrpc ApiGrcp.ApiGrpc
}

func (server *App) Run() {
	server.loadDotEnv()
	db := server.loadDb()
	server.loadLoop(db)
	server.loadApiGrpc(db)
}

func (server *App) loadDotEnv() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}

func (server *App) loadDb() *gorm.DB {
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
		log.Fatalf("failed to connect database %v", err)
	}
	return db
}

func (server *App) loadLoop(db *gorm.DB) {
	taskAdapter := Task.Adapter{Orm: db}
	resultAdapter := Result.Adapter{Orm: db}
	batchAdapter := Result.BatchAdapter{Orm: db}
	//TODO: inject grcp and proto? to test this adapter later
	callAdapter := Call.Adapter{}
	loopService := Loop.New(
		callAdapter,
		taskAdapter,
		taskAdapter,
		batchAdapter,
		resultAdapter,
	)
	loopManager := LoopManager.LoopManager{
		Loop: loopService,
	}
	go loopManager.Execute()
}

func (server *App) loadApiGrpc(db *gorm.DB) {
	var taskAdapter = Task.Adapter{Orm: db}
	var createTaskService = CreateTask.Service{SavePort: taskAdapter}
	var listTasksService = ListTasks.Service{FindByPort: taskAdapter}
	var showTaskService = ShowTask.Service{FindByPort: taskAdapter}
	server.ApiGrpc = ApiGrcp.ApiGrpc{}
	server.ApiGrpc.Initialize(
		createTaskService,
		listTasksService,
		showTaskService,
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	server.ApiGrpc.Run()

}
