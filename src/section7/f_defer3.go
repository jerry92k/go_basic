package main

import "fmt"

func stack() {

	for i := 1; i < 10; i++ {
		defer fmt.Println("ex 1 : ", i)
	}
}

func main() {

	stack()
}
