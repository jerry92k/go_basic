package main

import "fmt"

func start(t string) string {
	fmt.Println(" start: ", t)
	return t
}
func end(t string) {
	fmt.Println(" end : ", t)
}

func a() {
	defer end(start("b")) //중첩함수 주의. start 함수는 먼저 실행됨.
	fmt.Println("in a")
}

func main() {
	a()
}
