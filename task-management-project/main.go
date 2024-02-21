package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	name      string
	completed bool
}

var tasks = []Task{}

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

func main() {
	fmt.Println("Welcome to trello")

	for {

		reader := bufio.NewReader(os.Stdin)
		option, _ := getInput(reader, "Choose option (1 - create task, 2 - mark task as complete, 3 - delete tasks - 4 view tasks ): ")

		if option == "1" {
			createTask()
		}

		if option == "2" {
			completeTask()
		}

		if option == "3" {
			deleteTask()
		}

		if option == "4" {
			getAllTask()
		}

	}
}
