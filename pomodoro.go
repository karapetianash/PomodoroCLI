package main

import "time"

type Pomodoro struct {
	TaskName  string
	StartTime time.Time
}

func AddTask(taskName string) Pomodoro {
	return Pomodoro{
		TaskName:  taskName,
		StartTime: time.Now(),
	}
}

func TaskName(task Pomodoro) string {
	if taskIsEmpty(task) {
		return "Current task is empty."
	}

	return "Task: " + task.TaskName
}

func taskIsEmpty(task Pomodoro) bool {
	return Pomodoro{} == task
}
