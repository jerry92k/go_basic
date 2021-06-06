package main

import (
	"fmt"
)

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("hi")
	}()
	defer func() {
		fmt.Println("end?")
	}()
}

func main() {
	sayHello("Golang!")
}
