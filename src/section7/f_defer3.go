package main

import "fmt"

func stack() {

	for i := 1; i < 10; i++ {
		defer fmt.Println("ex 1 : ", i)
	}
}

func main() {

	//defer는 스택의 구조처럼 FILO 먼저 예약된게 나중에 실행됨
	stack()
}
