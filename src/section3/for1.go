package main

import "fmt"

func main() {
	// Go에서 유일하게 제공되는 반복문

	// 예제1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 에러발생 1. 여는 중괄호를 붙여 써야함
	// for i:=0; i<5; i++
	// {
	//
	// }

	// 에러발생 2
	/*
	  for i:=0; i<5; i++
	    fmt.Println(i)
	*/

	// 예제2. 무한루프
	for {
		fmt.Println("ex2 : Hello, Golang")
		fmt.Println("ex2 : Infinite Golang")
	}

	// 예제3. (Range용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println(index, name)
	}
}
