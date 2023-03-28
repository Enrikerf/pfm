package Looper

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Repository"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"os"
	"sync"
	"text/tabwriter"
)

var once sync.Once
var instance Looper

type Looper interface {
	IsEnabled() bool
	Enable()
}

func NewLooper() Looper {
	once.Do(func() {
		instance = &looper{
			isLoopEnabled: make(chan bool, 1),
		}
	})
	return instance
}

type looper struct {
	communicateRepository Repository.Communicate
	findTasksByRepository TaskRepository.FindBy
	saveTaskRepository    TaskRepository.Save
	saveBatchRepository   ResultRepository.SaveBatch
	saveResultRepository  ResultRepository.Save
	isLoopEnabled         chan bool
}

func (l *looper) IsEnabled() bool {
	return len(l.isLoopEnabled) != 0
}

func (l *looper) Enable() {
	l.isLoopEnabled <- true
	go l.loop()
}

func (l *looper) loop() {
	for l.IsEnabled() {
		tasks, err := l.findTasksByRepository.FindBy(map[string]interface{}{
			"Status":         Status.Pending,
			"execution_mode": ExecutionMode.Automatic,
		})
		if err != nil {
			fmt.Printf(err.Error())
			l.stopLoop()
		}
		if tasks == nil {
			fmt.Printf("No more task, loop stopped")
			l.stopLoop()
		}

		var wg sync.WaitGroup
		for index := range tasks {
			l.printTask(index, tasks[index])
			wg.Add(1)
			go l.executeTask(&wg, tasks[index])
		}
		wg.Wait()
	}
}

func (l *looper) executeTask(wg *sync.WaitGroup, task Task.Task) {
	defer wg.Done()
	task.SetStatus(Status.New(Status.Running))
	l.saveTaskRepository.Persist(task)
	resultBatch := l.communicateRepository.Communicate(task)
	l.saveBatchRepository.Persist(resultBatch)
	task.SetStatus(Status.New(Status.Done))
	l.saveTaskRepository.Persist(task)
}

func (l *looper) stopLoop() {
	if len(l.isLoopEnabled) > 0 {
		fmt.Println("loop stopped")
		<-l.isLoopEnabled
	} else {
		fmt.Println("trying to stop loop but is not running")
	}
}

func (l *looper) printTask(index int, task Task.Task) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	_, err := fmt.Fprintf(w, "%v) task:  \n"+
		"\t uuid \t host \t port \t command \t mode \t status\n"+
		"\t %v \t %v \t %v \t %v \t %v \t %v\n",
		index,
		task.GetId().GetUuidString(),
		task.GetHost().GetValue(),
		task.GetPort().GetValue(),
		len(task.GetSteps()),
		task.GetExecutionMode(),
		task.GetStatus().Value(),
	)
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}
	fmt.Println("----")
}
