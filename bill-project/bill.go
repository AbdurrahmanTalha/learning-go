package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func createBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.00,
	}

	return b
}

func (b bill) getBill() {
	fmt.Printf("Testing")
}
