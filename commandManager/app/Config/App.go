package Config

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/LoopManager"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Grcp/Call"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Adapters"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/CreateBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/DeleteBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ListBatches"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ReadBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/UpdateBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Call/Loop"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/CreateResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/DeleteResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ListResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ReadResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/UpdateResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/CreateStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/DeleteStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/ListSteps"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/ReadStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/UpdateStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ReadTask"
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
	taskAdapter := Adapters.TaskAdapter{Orm: db}
	resultAdapter := Adapters.ResultAdapter{Orm: db}
	batchAdapter := Adapters.BatchAdapter{Orm: db}
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
	var taskAdapter = Adapters.TaskAdapter{Orm: db}
	var commandAdapter = Adapters.StepAdapter{Orm: db}
	var batchAdapter = Adapters.BatchAdapter{Orm: db}
	var resultAdapter = Adapters.ResultAdapter{Orm: db}

	//TaskController
	var createTaskService = CreateTask.Service{SaveTaskPort: taskAdapter}
	var readTaskService = ReadTask.Service{FindTaskPort: taskAdapter}
	var updateTaskService = UpdateTask.Service{
		FindTaskPort:   taskAdapter,
		UpdateTaskPort: taskAdapter,
	}
	var deleteTaskService = DeleteTask.Service{DeleteTaskPort: taskAdapter}
	var listTasksService = ListTasks.Service{FindTasksByPort: taskAdapter}

	// CommandController
	var createStepService = CreateStep.Service{
		FindTaskPort: taskAdapter,
		SaveStepPort: commandAdapter,
	}
	var readStepService = ReadStep.Service{FindStepPort: commandAdapter}
	var updateStepService = UpdateStep.Service{
		FindStepPort:   commandAdapter,
		FindTaskPort:   taskAdapter,
		UpdateStepPort: commandAdapter,
	}
	var deleteStepService = DeleteStep.Service{DeleteStepPort: commandAdapter}
	var listStepsService = ListSteps.Service{FindStepByPort: commandAdapter}

	// BatchController
	var createBatchService = CreateBatch.Service{SaveBatchPort: batchAdapter}
	var readBatchService = ReadBatch.Service{FindBatchPort: batchAdapter}
	var updateBatchService = UpdateBatch.Service{
		FindBatchPort:   batchAdapter,
		FindTaskPort:    taskAdapter,
		UpdateBatchPort: batchAdapter,
	}
	var deleteBatchService = DeleteBatch.Service{DeleteBatchPort: batchAdapter}
	var listBatchesService = ListBatches.Service{FindBatchesByPort: batchAdapter}

	//ResultController
	var createResultService = CreateResult.Service{
		FindBatchPort:  batchAdapter,
		SaveResultPort: resultAdapter,
	}
	var readResultService = ReadResult.Service{FindResultPort: resultAdapter}
	var updateResultService = UpdateResult.Service{
		FindResultPort:   resultAdapter,
		FindBatchPort:    batchAdapter,
		UpdateResultPort: resultAdapter,
	}
	var deleteResultService = DeleteResult.Service{DeleteTaskPort: resultAdapter}
	var listResultsService = ListResults.Service{FindResultsByPort: resultAdapter}

	server.ApiGrpc = ApiGrcp.ApiGrpc{}
	server.ApiGrpc.Initialize(
		createTaskService,
		readTaskService,
		updateTaskService,
		deleteTaskService,
		listTasksService,
		createResultService,
		readResultService,
		updateResultService,
		deleteResultService,
		listResultsService,
		createBatchService,
		readBatchService,
		updateBatchService,
		deleteBatchService,
		listBatchesService,
		createStepService,
		readStepService,
		updateStepService,
		deleteStepService,
		listStepsService,
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	server.ApiGrpc.Run()

}
