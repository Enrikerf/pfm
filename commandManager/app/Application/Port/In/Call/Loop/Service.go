package Loop

import (
	"fmt"
	ResultOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	CallOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Grcp/Call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

// Manager TODO: rename Service
//TODO: construct throw new to make variables unexported
type Manager struct {
	CallRequestPort CallOutPort.RequestPort
	FindTasksByPort TaskOutPort.FindByPort
	UpdateTaskPort  TaskOutPort.UpdatePort
	SaveResultPort  ResultOutPort.SavePort
}

func (manager Manager) Loop() error {

	for true {
		manager.Iteration()
		time.Sleep(10 * time.Second)
	}
	return nil
}

func (manager Manager) Iteration() {
	fmt.Println("--- loop ----")
	tasks, err := manager.FindTasksByPort.FindBy(map[string]interface{}{"Status": TaskDomain.Pending.String()})
	if err != nil {
		fmt.Printf("error fetching task: %v. \n", err)
	}
	fmt.Printf("tasks %v. \n", len(tasks))
	var wg sync.WaitGroup
	for index := range tasks {
		wg.Add(1)
		go manager.slot(&wg, index, &tasks[index])
	}
	wg.Wait()
}

func (manager Manager) slot(wg *sync.WaitGroup, index int, tasks *TaskDomain.Task) {
	defer wg.Done()
	printTask(index, *tasks)
	manager.updateTaskStatus(tasks, TaskDomain.Running)
	result := manager.callToClient(tasks)
	manager.saveResult(result)
	manager.updateTaskStatus(tasks, TaskDomain.Done)
	println(tasks.Status.String())
}

func (manager Manager) updateTaskStatus(task *TaskDomain.Task, status TaskDomain.TaskStatus) {
	task.Status = status
	err := manager.UpdateTaskPort.Update(*task)
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

func (manager Manager) callToClient(task *TaskDomain.Task) Result.Result {
	return manager.CallRequestPort.Request(*task)
}

//func (manager Manager) callToClient(task TaskDomain.Task) Result.Result {
//	resultChannel := make(chan Result.Result)
//	go func(task TaskDomain.Task) {
//		defer close(resultChannel)
//		resultChannel <- manager.CallRequestPort.Request(task)
//	}(task)
//	result := <-resultChannel
//	return result
//}

func printTask(index int, task TaskDomain.Task) {
	fmt.Printf("%v) task:  \n", index)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	fmt.Fprintln(w, "uuid \t host \t port \t command \t mode \t status")
	fmt.Fprintf(w, "%v \t %v \t %v \t %v \t %v \t %v\n", task.Uuid, task.Host, task.Port, task.Command, task.Mode, task.Status)
	w.Flush()
}
