// 함수 closure

package main

import "fmt"

func main() {
	// 클로저
	// 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	// 함수 안에서 함수를 선언 및 정의 가능 -> 이때 외부 함수에 선언된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)

	// 예제1

	multiply := func(x, y int) int { //익명함수
		return x * y
	}
	r1 := multiply(10, 5)

	fmt.Println(r1)

	m, n := 5, 10
	sum := func(c int) int {
		return m + n + c
	}
	r2 := sum(10)

	fmt.Println(r2)
}
