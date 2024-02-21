package main

import (
	"fmt"
)

type dojo struct {
	name     string
	master   *ninja
	students []*ninja
}

type ninja struct {
	name    string
	Weapons []string `default:"[\"nothing\"]"`
	level   int
}

func (n ninja) sayHello() {
	fmt.Println("Hello I'm", n.name, "and i'm on level", n.level)

}

func main() {
	masterNinja := ninja{"master", []string{"EVERYTHING"}, 5}
	masterDojo := dojo{
		name:     "The Master Dojo",
		master:   &masterNinja,
		students: []*ninja{},
	}
	masterDojo.master.sayHello()
}

// * A project where youtube is in focus mode
