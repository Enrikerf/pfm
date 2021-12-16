package LoopManager

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Task"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

type LoopManager struct {
	DB *gorm.DB
}

func (loopManager LoopManager) Run() {

	//for true {
	fmt.Println("exec")
	var tasks []Task.Task
	response := loopManager.DB.Find(&tasks)
	fmt.Printf("tasks %v. \n", response.RowsAffected)
	if response.Error != nil {
		fmt.Printf("tasks %v. \n", response.Error)
	}
	for index, task := range tasks {
		printTask(index, task)
		response := callToClient(task)
		saveResult(response)
		loopManager.updateTaskSTatus(task)
	}
	time.Sleep(10 * time.Second)
	//}
}

func (loopManager LoopManager) updateTaskSTatus(task Task.Task) {
	task.Status = TaskDomain.Done.String()
	loopManager.DB.Save(&task)
}

func saveResult(result string) {
	println("saving result in db: " + result)
}

func callToClient(task Task.Task) string {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.Host+":"+task.Port, options)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer connection.Close()
	client := call.NewCallServiceClient(connection)
	callRequest := call.CallRequest{
		Command: task.Command,
	}
	var callResponse2 string
	switch task.Mode {
	case TaskDomain.Unary.String():
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
		}
		callResponse2 = callResponse.GetResult()
	case TaskDomain.ServerStream.String():
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
		}
		callResponse2 = callResponse.GetResult()
	case TaskDomain.ClientStream.String():
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
		}
		callResponse2 = callResponse.GetResult()
	case TaskDomain.Bidirectional.String():
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
		}
		callResponse2 = callResponse.GetResult()
	default:
		callResponse2 = "not implemented task mode"
		fmt.Printf("%v) task:  \n", err)
	}

	return callResponse2
}

func printTask(index int, task Task.Task) {
	fmt.Printf("%v) task:  \n", index)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	fmt.Fprintln(w, "uuid \t host \t port \t command \t mode \t status")
	fmt.Fprintf(w, "%v \t %v \t %v \t %v \t %v \t %v\n", task.ID, task.Host, task.Port, task.Command, task.Mode, task.Status)
	w.Flush()
}
