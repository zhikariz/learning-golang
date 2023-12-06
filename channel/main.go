package main

import "fmt"

func main() {
	var message = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		message <- data
	}

	go sayHelloTo("John")
	go sayHelloTo("Wick")
	go sayHelloTo("Jason Bourne")

	var message1 = <-message
	fmt.Println(message1)

	var message2 = <-message
	fmt.Println(message2)

	var message3 = <-message
	fmt.Println(message3)
}
