package Loop

import (
	"fmt"
	ResultOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Result"
	TaskOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Database/Task"
	CallOutPort "github.com/Enrikerf/pfm/commandManager/app/Application/Port/Out/Grcp/Call"
	ResultDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Result"
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Model/Task"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

type Service struct {
	callRequestPort CallOutPort.RequestPort
	findTasksByPort TaskOutPort.FindByPort
	updateTaskPort  TaskOutPort.UpdatePort
	saveBatchPort   ResultOutPort.SaveBatchPort
	saveResultPort  ResultOutPort.SaveResultPort
	exit            bool
}

func New(
	callRequestPort CallOutPort.RequestPort,
	findTasksByPort TaskOutPort.FindByPort,
	updateTaskPort TaskOutPort.UpdatePort,
	saveBatchPort ResultOutPort.SaveBatchPort,
	saveResultPort ResultOutPort.SaveResultPort) *Service {
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
	fmt.Println("------------ LOOP -------------")
	tasks := service.findTasksByPort.
		FindBy(map[string]interface{}{"Status": TaskDomain.Pending.String(), "execution_mode": TaskDomain.Automatic})
	if tasks == nil {
		fmt.Printf("error fetching task from db. \n")
		service.exit = true
	} else {
		fmt.Printf("tasks %v. \n", len(tasks))
		var wg sync.WaitGroup
		for index := range tasks {
			service.printTask(index, tasks[index])
			wg.Add(1)
			go service.slot(&wg, index, &tasks[index])
		}
		wg.Wait()
	}

}

func (service *Service) slot(wg *sync.WaitGroup, index int, task *TaskDomain.Task) {
	defer wg.Done()
	service.updateTaskStatus(index, task, TaskDomain.Running)
	resultBatch := service.callRequestPort.Request(*task)
	service.saveResults(index, resultBatch)
	service.updateTaskStatus(index, task, TaskDomain.Done)
}

func (service *Service) updateTaskStatus(index int, task *TaskDomain.Task, status TaskDomain.TaskStatus) {
	task.Status = status
	err := service.updateTaskPort.Update(*task)
	if err != nil {
		fmt.Printf("\t%v-error updating task %v. \n", index, err)
		service.exit = true
	}
}

func (service *Service) saveResults(index int, resultBatch ResultDomain.Batch) {
	fmt.Printf("\t%v-saving result in db: %v\n", index, resultBatch)
	err := service.saveBatchPort.Save(resultBatch)
	if err != nil {
		fmt.Printf("\t%v-error saving result %v. \n", index, err)
		service.exit = true
	}
	//for index := range resultBatch {
	//	fmt.Printf("\t%v-saving result in db: %v\n", index, resultBatch[index].Content)
	//	err := service.saveResultPort.Save(resultBatch[index])
	//	if err != nil {
	//		fmt.Printf("\t%v-error saving result %v. \n", index, err)
	//		service.exit = true
	//	}
	//}

}

func (service *Service) printTask(index int, task TaskDomain.Task) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	_, err := fmt.Fprintf(w, "%v) task:  \n"+
		"\t uuid \t host \t port \t command \t mode \t status\n"+
		"\t %v \t %v \t %v \t %v \t %v \t %v\n",
		index, task.Uuid, task.Host, task.Port, task.Commands, task.Mode, task.Status)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}
	fmt.Println("----")
}
