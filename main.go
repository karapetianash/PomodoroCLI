package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var currentTask Pomodoro = Pomodoro{}

func main() {
	taskPtr := flag.String("task", "", "Task to add. (Required)")
	flag.Parse()

	if *taskPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	currentTask = AddTask(*taskPtr)
	fmt.Printf("Start the task! Focus on %s\n", *taskPtr)

	timer := time.NewTimer(5 * time.Second)
	<-timer.C

	fmt.Println("Congrats! Task time is complete. Take a break.")
}
