package main

import (
	"fmt"
	d1 "golang-training/day_1"
	"os"
)

func main() {
	defer func() {
		os.Exit(0)
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	d1.Day1()
	panic("abdljf")
}
