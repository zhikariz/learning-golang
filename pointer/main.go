package main

import "fmt"

func main() {
	numberA := 4
	var numberB *int = &numberA
	checkResult(numberA, numberB)
	fmt.Println("============================")
	changeValue(numberB, 5)
	checkResult(numberA, numberB)
}

func checkResult(numberA int, numberB *int) {
	fmt.Println("numberA (value)", numberA)
	fmt.Println("numberA (address)", &numberA)

	fmt.Println("numberB (value)", *numberB)
	fmt.Println("numberB (address)", numberB)
}

func changeValue(original *int, value int) {
	*original = value
}
