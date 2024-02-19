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

/* for i := 0; i < 10; i++ {
	fmt.Println(i)
}

x := 0
for x < 20 {
	fmt.Println(x)
	x++
} */

// names := []string{"mario", "luigi"}

/* for i := 0; i < len(names); i++ {
	fmt.Println(names[i])
} */

/* for index, value := range names {
	fmt.Println(index, value)
} */

/* package main

import (
	"fmt"
)

type Item struct {
	Name string
	Price float32
}

func main() {
	var billName string;
	var items []Item

	fmt.Print("Create a new bill: ");
	fmt.Scan(&billName);
	fmt.Println("Successfully created new bill", billName);

	for {
		var option string;
		fmt.Print("Choose option (a - add item, s - save bill, t - add tip): ")
		fmt.Scanln(&option);
		if option == "a" {
			var itemName string;
			var itemPrice float32;

			fmt.Printf("Item name: ");
			fmt.Scanln(&itemName);
			fmt.Printf("Item Price: ")
			fmt.Scanln(&itemPrice);

			items = append(items, Item{Name: itemName, Price: itemPrice})
			fmt.Println("Successfully added item")
		} else if option == "s" {
			fmt.Println(items);
			break;
		} else if option == "t" {
			var tipAmount float32;
			fmt.Printf("Enter tip amount ($): ")
			fmt.Scanln(&tipAmount);

			items = append(items, Item{Name: "Tip", Price: tipAmount})
		} else {
			fmt.Println("Invalid command try again")
		}
	}

}
*/
