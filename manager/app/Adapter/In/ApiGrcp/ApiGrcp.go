package ApiGrcp

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/Controller"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/batch"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/result"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/step"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/CreateBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/DeleteBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ListBatches"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/ReadBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Batch/UpdateBatch"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/CreateResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/DeleteResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ListResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/ReadResult"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/StreamResults"
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
	readTaskUseCase   ReadTask.UseCase
	deleteTaskUseCase DeleteTask.UseCase
	updateTaskUseCase UpdateTask.UseCase

	createResultUseCase  CreateResult.UseCase
	readResultUseCase    ReadResult.UseCase
	updateResultUseCase  UpdateResult.UseCase
	deleteResultUseCase  DeleteResult.UseCase
	listResultsUseCase   ListResults.UseCase
	streamResultsUseCase StreamResults.UseCase

	createBatchUseCase CreateBatch.UseCase
	readBatchUseCase   ReadBatch.UseCase
	updateBatchUseCase UpdateBatch.UseCase
	deleteBatchUseCase DeleteBatch.UseCase
	listBatchesUseCase ListBatches.UseCase

	createStepUseCase CreateStep.UseCase
	readStepUseCase   ReadStep.UseCase
	updateStepUseCase UpdateStep.UseCase
	deleteStepUseCase DeleteStep.UseCase
	listStepsUseCase  ListSteps.UseCase

	serverHost string
	serverPort string
	grpcServer *grpc.Server
	listener   net.Listener
}

func (api *ApiGrpc) Initialize(
	createTaskUseCase CreateTask.UseCase,
	readTaskUseCase ReadTask.UseCase,
	updateTaskUseCase UpdateTask.UseCase,
	deleteTaskUseCase DeleteTask.UseCase,
	listTasksUseCase ListTasks.UseCase,

	createResultUseCase CreateResult.UseCase,
	readResultUseCase ReadResult.UseCase,
	updateResultUseCase UpdateResult.UseCase,
	deleteResultUseCase DeleteResult.UseCase,
	listResultsUseCase ListResults.UseCase,
	streamResultsUseCase StreamResults.UseCase,

	createBatchUseCase CreateBatch.UseCase,
	readBatchUseCase ReadBatch.UseCase,
	updateBatchUseCase UpdateBatch.UseCase,
	deleteBatchUseCase DeleteBatch.UseCase,
	listBatchesUseCase ListBatches.UseCase,

	createStepUseCase CreateStep.UseCase,
	readStepUseCase ReadStep.UseCase,
	updateStepUseCase UpdateStep.UseCase,
	deleteStepUseCase DeleteStep.UseCase,
	listStepsUseCase ListSteps.UseCase,

	host string,
	port string,
) {
	fmt.Println("Starting Sentence Manager...")

	api.createTaskUseCase = createTaskUseCase
	api.listTasksUseCase = listTasksUseCase
	api.readTaskUseCase = readTaskUseCase
	api.deleteTaskUseCase = deleteTaskUseCase
	api.updateTaskUseCase = updateTaskUseCase

	api.listResultsUseCase = listResultsUseCase
	api.createResultUseCase = createResultUseCase
	api.readResultUseCase = readResultUseCase
	api.updateResultUseCase = updateResultUseCase
	api.deleteResultUseCase = deleteResultUseCase
	api.listResultsUseCase = listResultsUseCase
	api.streamResultsUseCase = streamResultsUseCase

	api.createBatchUseCase = createBatchUseCase
	api.readBatchUseCase = readBatchUseCase
	api.updateBatchUseCase = updateBatchUseCase
	api.deleteBatchUseCase = deleteBatchUseCase
	api.listBatchesUseCase = listBatchesUseCase

	api.createStepUseCase = createStepUseCase
	api.readStepUseCase = readStepUseCase
	api.updateStepUseCase = updateStepUseCase
	api.deleteStepUseCase = deleteStepUseCase
	api.listStepsUseCase = listStepsUseCase

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
			log.Fatalf(err.Error())
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
		ReadTaskUseCase:   api.readTaskUseCase,
		DeleteTaskUseCase: api.deleteTaskUseCase,
		UpdateTaskUseCase: api.updateTaskUseCase,
	}
	task.RegisterTaskServiceServer(api.grpcServer, taskController)

	var resultController = Controller.ResultController{
		CreateResultUseCase:  api.createResultUseCase,
		ReadResultUseCase:    api.readResultUseCase,
		UpdateResultUseCase:  api.updateResultUseCase,
		DeleteResultUseCase:  api.deleteResultUseCase,
		ListResultsUseCase:   api.listResultsUseCase,
		StreamResultsUseCase: api.streamResultsUseCase,
	}
	result.RegisterResultServiceServer(api.grpcServer, resultController)

	var batchController = Controller.BatchController{
		CreateBatchUseCase: api.createBatchUseCase,
		ReadBatchUseCase:   api.readBatchUseCase,
		UpdateBatchUseCase: api.updateBatchUseCase,
		DeleteBatchUseCase: api.deleteBatchUseCase,
		ListBatchesUseCase: api.listBatchesUseCase,
	}
	batch.RegisterBatchServiceServer(api.grpcServer, batchController)

	var commandController = Controller.StepController{
		CreateStepUseCase: api.createStepUseCase,
		ReadStepUseCase:   api.readStepUseCase,
		UpdateStepUseCase: api.updateStepUseCase,
		DeleteStepUseCase: api.deleteStepUseCase,
		ListStepsUseCase:  api.listStepsUseCase,
	}
	step.RegisterStepServiceServer(api.grpcServer, commandController)
}

func (api *ApiGrpc) loadServer() {
	//var serverOptions []grpc.ServerOption
	errHandler := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			fmt.Printf("method %q failed: %s", info.FullMethod, err)
		}
		return resp, err
	}
	api.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(errHandler))
	if os.Getenv("APP_DEBUG") == "true" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

func (api *ApiGrpc) loadListener() {
	listener, err := net.Listen("tcp", api.serverHost+api.serverPort)
	if err != nil {
		//log.Fatalf("failed to listen at: " + api.serverHost + api.serverPort)
		log.Fatalf(err.Error())
	}
	api.listener = listener
}
