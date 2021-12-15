package Config

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type App struct {
	ApiGrpc ApiGrcp.ApiGrpc
}

func (server *App) Run() {

	loadDotEnv()
	server.ApiGrpc = ApiGrcp.ApiGrpc{}
	db := loadDb()
	server.ApiGrpc.Initialize(
		db,
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	go commandManagerLoop(db)
	server.ApiGrpc.Run()
}

func commandManagerLoop(db *gorm.DB) {
	for true {
		// TODO: wrap in layers and develop logic
		fmt.Println("exec")
		var tasks []Task.Task
		result:=db.Find(&tasks)
		fmt.Printf("tasks %v",result.RowsAffected)
		time.Sleep(10* time.Second)
	}
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
