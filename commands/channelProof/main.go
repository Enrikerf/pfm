package main

import (
	"fmt"
	"time"
)

func main() {

	var i int
	channel := make(chan bool, 1)
	for {
		fmt.Println("commands:")
		fmt.Println("\t1) controller up")
		fmt.Println("\t2) controller down")
		fmt.Printf("choose: ")
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		switch i {
		case 1:
			go controller(channel)
		case 2:
			if len(channel) > 0 {
				println("clossing")
				<-channel
			} else {
				println("no proccess running")
			}
		default:
			fmt.Println("- command not found")
		}
	}

}

func controller(channel chan bool) {
	sampleTime := time.Second * 1
	channel <- true
	for range time.Tick(sampleTime) {
		fmt.Println("\nhello")
		if len(channel) == 0 {
			fmt.Println("\nbye")
			break
		}
	}
}
