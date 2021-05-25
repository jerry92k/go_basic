package main

import (
	"fmt"
)

// init 이 여러개여도 상관없음

func init() {
	fmt.Println("init 1")
}
func init() {
	fmt.Println("init 2")
}
func init() {
	fmt.Println("init 3")
}
func init() {
	fmt.Println("init 4")
}

func main() {

	fmt.Println("main method start")
}
