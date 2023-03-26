package Config

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/InOut/EventDispatcherAdapter"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Grcp/Call"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Adapters/ResultAdapter"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Adapters/TaskAdapter"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/CreateBatchAndFill"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/FillBatchEventHandler"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/GetBatchResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/GetTaskBatches"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/StreamResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ReadTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/TaskEventHandler"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/UpdateTask"
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
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}

func (server *App) loadApiGrpc(db *gorm.DB) {

	var findTaskAdapter = TaskAdapter.FindAdapter{Orm: db}
	var deleteAdapter = TaskAdapter.DeleteAdapter{Orm: db}
	var saveTaskAdapter = TaskAdapter.PersistAdapter{Orm: db}
	var findTasksByAdapter = TaskAdapter.FindByAdapter{Orm: db}
	var findBatchAdapter = ResultAdapter.FindBatchAdapter{Orm: db}
	var saveBatchAdapter = ResultAdapter.PersistBatchAdapter{Orm: db}
	var saveResultAdapter = ResultAdapter.PersistAdapter{Orm: db}
	var findBatchResultsAdapter = ResultAdapter.FindBatchResultsAdapter{Orm: db}
	var findTaskBatchesAdapter = ResultAdapter.FindTaskBatchesAdapter{Orm: db}
	var findBatchResultsAfterResultAdapter = ResultAdapter.FindBatchResultsAfterResultAdapter{Orm: db}
	var callAdapter = Call.New()

	//Task UseCases
	var taskEventHandler = TaskEventHandler.New(findTaskAdapter)
	var fillBatchEventHandler = FillBatchEventHandler.New(callAdapter, findBatchAdapter, findTaskAdapter, saveResultAdapter)
	//EventDispatcher
	var eventDispatcherAdapter = EventDispatcherAdapter.New(taskEventHandler, fillBatchEventHandler)
	//--------------
	var createTaskService = CreateTask.New(saveTaskAdapter, eventDispatcherAdapter)
	var readTaskService = ReadTask.New(findTaskAdapter)
	var updateTaskService = UpdateTask.New(findTaskAdapter, saveTaskAdapter, eventDispatcherAdapter)
	var deleteTaskService = DeleteTask.New(findTaskAdapter, deleteAdapter)
	var listTasksService = ListTasks.New(findTasksByAdapter)
	//Result UseCase
	var executeTaskManually = CreateBatchAndFill.New(
		eventDispatcherAdapter,
		findTaskAdapter,
		saveTaskAdapter,
		saveBatchAdapter,
	)
	var listResultsService = GetBatchResults.New(findBatchResultsAdapter)
	var getTaskBatches = GetTaskBatches.New(findTaskBatchesAdapter)
	var streamResults = StreamResults.New(findBatchAdapter, findTaskAdapter, findBatchResultsAfterResultAdapter)

	server.ApiGrpc = ApiGrcp.ApiGrpc{}
	server.ApiGrpc.Initialize(
		createTaskService,
		readTaskService,
		updateTaskService,
		deleteTaskService,
		listTasksService,
		executeTaskManually,
		listResultsService,
		streamResults,
		getTaskBatches,
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	server.ApiGrpc.Run()

}
