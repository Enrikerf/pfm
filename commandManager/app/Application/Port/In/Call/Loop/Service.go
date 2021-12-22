package Loop

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	ResultOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

type Manager struct {
	FindAllTaskPort TaskOutPort.FindAllPort
	SaveTaskPort   TaskOutPort.SavePort
	SaveResultPort ResultOutPort.SavePort
}

func (manager Manager) Loop() {

	//for true {
	fmt.Println("exec")
	tasks,err := manager.FindAllTaskPort.FindAll()
	fmt.Printf("tasks %v. \n", len(tasks))
	if err != nil {
		fmt.Printf("tasks %v. \n", err)
	}
	for index, task := range tasks {
		printTask(index, task)
		response := manager.callToClient(task)
		manager.saveResult(response)
		manager.updateTaskSTatus(task)
	}
	time.Sleep(10 * time.Second)
	//}
}

func (manager Manager) updateTaskSTatus(task TaskDomain.Task) {
	task.Status = TaskDomain.Done
	manager.SaveTaskPort.Save(task)
}

func (manager Manager) saveResult(result string) {
	println("saving result in db: " + result)
	_ = manager.SaveResultPort.Save(Result.Result{
		Id:      uuid.UUID{},
		TaskId:  uuid.UUID{},
		Content: "fake content",
	})
}

func (manager Manager) callToClient(task TaskDomain.Task) string {
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
	case TaskDomain.Unary:
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
		}
		callResponse2 = callResponse.GetResult()
	case TaskDomain.ServerStream:
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
		}
		callResponse2 = callResponse.GetResult()
	case TaskDomain.ClientStream:
		callResponse, err := client.CallUnary(context.Background(), &callRequest)
		if err != nil {
			fmt.Printf("%v) task:  \n", err)
		}
		callResponse2 = callResponse.GetResult()
	case TaskDomain.Bidirectional:
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

func printTask(index int, task TaskDomain.Task) {
	fmt.Printf("%v) task:  \n", index)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	fmt.Fprintln(w, "uuid \t host \t port \t command \t mode \t status")
	fmt.Fprintf(w, "%v \t %v \t %v \t %v \t %v \t %v\n", task.Id, task.Host, task.Port, task.Command, task.Mode, task.Status)
	w.Flush()
}