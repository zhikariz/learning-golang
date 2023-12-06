package main

import "fmt"

func main() {
	var name string
	fmt.Println("Type your name")
	fmt.Scanln(&name)

	defer catch()

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		panic(err.Error())
	}
}

func validate(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}
	if len(name) > 10 {
		return false, fmt.Errorf("name cannot be more than 10 characters")
	}
	return true, nil
}

func catch() {
	if err := recover(); err != nil {
		fmt.Println("error occured", err)
	}
}
