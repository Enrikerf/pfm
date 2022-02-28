package Loop

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/BatchPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/ResultPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/TaskPort"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Grcp/CallPort"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Entity"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/ValueObject"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

type Service struct {
	callRequestPort CallPort.Request
	findTasksByPort TaskPort.FindBy
	updateTaskPort  TaskPort.Update
	saveBatchPort   BatchPort.Save
	saveResultPort  ResultPort.Save
	exit            bool
}

func New(
	callRequestPort CallPort.Request,
	findTasksByPort TaskPort.FindBy,
	updateTaskPort TaskPort.Update,
	saveBatchPort BatchPort.Save,
	saveResultPort ResultPort.Save) *Service {
	return &Service{
		callRequestPort: callRequestPort,
		findTasksByPort: findTasksByPort,
		updateTaskPort:  updateTaskPort,
		saveBatchPort:   saveBatchPort,
		saveResultPort:  saveResultPort,
		exit:            false,
	}
}

func (service *Service) GetExit() bool {
	return service.exit

}
func (service *Service) Loop() error {

	for !service.exit {
		service.Iteration()
		time.Sleep(10 * time.Second)
	}
	return nil
}

func (service *Service) Iteration() {
	//fmt.Println("------------ LOOP -------------")
	tasks := service.findTasksByPort.
		FindBy(map[string]interface{}{"Status": ValueObject.Pending.String(), "execution_mode": ValueObject.Automatic.String()})
	if tasks == nil {
		fmt.Printf("error fetching task from db. \n")
		service.exit = true
	} else {
		//fmt.Printf("tasks %v. \n", len(tasks))
		var wg sync.WaitGroup
		for index := range tasks {
			service.printTask(index, tasks[index])
			wg.Add(1)
			go service.slot(&wg, index, &tasks[index])
		}
		wg.Wait()
	}

}

func (service *Service) slot(wg *sync.WaitGroup, index int, task *Entity.Task) {
	defer wg.Done()
	service.ExecuteTask(index, task)
}

func (service *Service) ExecuteTask(index int, task *Entity.Task) {
	service.updateTaskStatus(index, task, ValueObject.Running)
	resultBatch := service.callRequestPort.Request(*task)
	service.saveResults(index, resultBatch)
	service.updateTaskStatus(index, task, ValueObject.Done)
}

func (service *Service) updateTaskStatus(index int, task *Entity.Task, status ValueObject.TaskStatus) {
	task.Status = status
	err := service.updateTaskPort.Update(*task)
	if err != nil {
		fmt.Printf("\t%v-error updating task %v. \n", index, err)
		service.exit = true
	}
}

func (service *Service) saveResults(index int, resultBatch Entity.Batch) {
	fmt.Printf("\t%v-saving result in db: %v\n", index, resultBatch)
	err := service.saveBatchPort.Save(resultBatch)
	if err != nil {
		fmt.Printf("\t%v-error saving result %v. \n", index, err)
		service.exit = true
	}
}

func (service *Service) printTask(index int, task Entity.Task) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	_, err := fmt.Fprintf(w, "%v) task:  \n"+
		"\t uuid \t host \t port \t command \t mode \t status\n"+
		"\t %v \t %v \t %v \t %v \t %v \t %v\n",
		index, task.Uuid, task.Host, task.Port, task.Steps, task.Mode, task.Status)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}
	fmt.Println("----")
}
