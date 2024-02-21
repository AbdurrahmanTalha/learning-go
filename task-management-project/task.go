package main

import (
	"bufio"
	"fmt"
	"os"
)

func createTask() {
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := getInput(reader, "Task name: ")

	tasks = append(tasks, Task{name: taskName, completed: false})
	fmt.Println("Successfully created task")
}

func getAllTask() {
	fmt.Println("Your todo tasks are: ")
	for _, value := range tasks {
		fmt.Printf("Name: %v; Completed: %t\n", value.name, value.completed)
	}
}

func completeTask() {
	var success bool
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := getInput(reader, "Task name: ")

	for i := range tasks {
		if tasks[i].name == taskName {
			tasks[i].completed = true
			success = true
			fmt.Println("Successfully marked task as successful")
			break
		}
	}
	if !success {
		fmt.Println("Failed to find task")
	}
}

func deleteTask() {
	var index int = -1
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := getInput(reader, "Task name: ")

	fmt.Println(taskName)

	for i := range tasks {
		if tasks[i].name == taskName {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Failed to find task")
	} else {
		tasks = append(tasks[:index], tasks[index+1:]...)
		fmt.Println("Successfully deleted task")
	}
}
