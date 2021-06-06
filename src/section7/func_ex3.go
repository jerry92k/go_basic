package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// 재귀함수(recursion)
	// 프로그램이 보기 쉽고, 코드 간결, 오류수정 용이
	// 코드 이해 어렵고, 기억 공간을 많이 사용

	// 예제1
	x := fact(4)
	fmt.Println(x)

	//익명함수 => 즉시실행(선언과 동시에)

	func() {
		fmt.Println("ex 1 : Anoymous")
	}()

	msg := "helelo golang"

	func(m string) {
		fmt.Println("ex2 :", m)
	}(msg)

	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 20)

	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println(r)
}
