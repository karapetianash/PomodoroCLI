package main

import "testing"

func TestAddTest(t *testing.T) {
	expectedTaskName := "TestingTask"
	newTask := AddTask(expectedTaskName)

	if newTask.TaskName != expectedTaskName {
		t.Errorf("TestAddTask Failed - Expected: %s Got: %s", expectedTaskName, newTask.TaskName)
	}
}

func TestTaskName_WithTask(t *testing.T) {
	expectedTaskName := "TestingTask"
	newTask := AddTask("TestingTask")

	printOutput := TaskName(newTask)
	expectedPrintOutput := "Task: " + expectedTaskName

	if printOutput != expectedPrintOutput {
		t.Errorf("TestTaskName Failed - Expected: %s Got %s", expectedPrintOutput, printOutput)
	}
}

func TestTaskName_EmptyTask(t *testing.T) {
	printOutput := TaskName(Pomodoro{})
	expectedPrintOutput := "Current task is empty."

	if printOutput != expectedPrintOutput {
		t.Errorf("TestTaskName Failed - Expected: %s Got %s", expectedPrintOutput, printOutput)
	}
}
