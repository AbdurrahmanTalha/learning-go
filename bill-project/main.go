package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt) 
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new a bill name: ", reader)

	_, error := newBill(name)
	if error != nil {
		fmt.Println("Bad ")
	}
	// return b
}

func main() {
	// bill :=
	createBill()

	// fmt.Println(bill)
	// bill.updateTip(25.55)

	// fmt.Println(bill.format())
}
