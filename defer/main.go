package main

import "fmt"

func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(food string) {
	defer fmt.Println("I'm done eating", food)
	fmt.Println("I'm eating", food)
}
