package Loop

import (
	"fmt"
	ResultOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	CallOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Grcp/Call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"os"
	"text/tabwriter"
	"time"
)

type Manager struct {
	CallRequestPort CallOutPort.RequestPort
	FindByPort      TaskOutPort.FindByPort
	UpdatePort      TaskOutPort.UpdatePort
	SaveResultPort  ResultOutPort.SavePort
}

func (manager Manager) Loop() error {

	for true {
		fmt.Println("--- loop ----")
		tasks, err := manager.FindByPort.FindBy(map[string]interface{}{"Status": TaskDomain.Pending.String()})
		if err != nil {
			fmt.Printf("error fetching task: %v. \n", err)
		}
		fmt.Printf("tasks %v. \n", len(tasks))
		for index, task := range tasks {
			printTask(index, task)
			result := manager.callToClient(task)
			manager.saveResult(result)
			manager.updateTaskStatus(task)
		}
		time.Sleep(10 * time.Second)
	}
	return nil
}

func (manager Manager) updateTaskStatus(task TaskDomain.Task) {
	task.Status = TaskDomain.Done
	err := manager.UpdatePort.Update(task)
	if err != nil {
		fmt.Printf("error updating task %v. \n", err)
	}
}

func (manager Manager) saveResult(result Result.Result) {
	println("saving result in db: " + result.Content)
	err := manager.SaveResultPort.Save(result)
	if err != nil {
		fmt.Printf("error saving result %v. \n", err)
	}
}

func (manager Manager) callToClient(task TaskDomain.Task) Result.Result {
	return manager.CallRequestPort.Request(task)
}

func printTask(index int, task TaskDomain.Task) {
	fmt.Printf("%v) task:  \n", index)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	fmt.Fprintln(w, "uuid \t host \t port \t command \t mode \t status")
	fmt.Fprintf(w, "%v \t %v \t %v \t %v \t %v \t %v\n", task.Uuid, task.Host, task.Port, task.Command, task.Mode, task.Status)
	w.Flush()
}
