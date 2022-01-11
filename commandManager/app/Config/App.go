package Config

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/LoopManager"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Grcp/Call"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Call/Loop"
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

	loadDotEnv()
	db := loadDb()
	//loop := LoopManager.LoopManager{DB: db}
	//go loop.Run()
	loadLoop(db)
	server.ApiGrpc = ApiGrcp.ApiGrpc{}
	server.ApiGrpc.Initialize(
		db,
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	server.ApiGrpc.Run()
}

func loadDotEnv() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}

func loadDb() *gorm.DB {
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

func loadLoop(db *gorm.DB) {
	taskAdapter := Task.Adapter{Orm: db}
	resultAdapter := Result.Adapter{Orm: db}
	callAdapter := Call.Adapter{}
	loopService := Loop.Manager{
		CallRequestPort: callAdapter,
		FindByPort:      taskAdapter,
		UpdatePort:      taskAdapter,
		SaveResultPort:  resultAdapter,
	}
	loopManager := LoopManager.LoopManager{
		Loop: loopService,
	}
	go loopManager.Execute()
}
