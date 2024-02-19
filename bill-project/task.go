package main

import (
	"fmt"
)

func createTask() {
	var taskName string
	fmt.Printf("Task name: ")
	fmt.Scan(&taskName)
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
	var taskName string
	var success bool
	fmt.Printf("Task name: ")
	fmt.Scan(&taskName)

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
	var taskName string
	var index int = -1
	fmt.Printf("Task name: ")
	fmt.Scan(&taskName)

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
