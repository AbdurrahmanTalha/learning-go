package main

import (
	"fmt"
)

type Task struct {
	name      string
	completed bool
}

var tasks = []Task{}

func main() {
	fmt.Println("Welcome to trello")
	for true {
		var option int

		fmt.Printf("Choose option (1 - create task, 2 - mark task as complete, 3 - delete tasks - 4 view tasks ): ")
		fmt.Scan(&option)

		if option == 1 {
			createTask()
		}

		if option == 2 {
			completeTask()
		}

		if option == 3 {
			deleteTask()
		}

		if option == 4 {
			getAllTask() 
		}

	}
}

