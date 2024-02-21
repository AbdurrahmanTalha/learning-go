package main

import (
	"errors"
	"fmt"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func newBill(name string) (Bill, error) {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.00,
	}

	return b, errors.New("Something went wrong!")
}

func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b Bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v  ...$%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-15v ...$%0.2f\n", "Tip:", b.tip)
	total += b.tip
	fs += fmt.Sprintf("%-15v ...$%0.2f", "Total:", total)

	return fs
}

func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}
