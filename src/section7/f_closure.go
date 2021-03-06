// 함수 closure

package main

import "fmt"

func main() {
	// 클로저
	// 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	// 함수 안에서 함수를 선언 및 정의 가능 -> 이때 외부 함수에 선언된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 함수를 호출할때 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트, 무분별한 전역변수 남발 줄
	// 남발 -> 메모리누수
	// 클로저를 정확하게 이해하고 사용해야함

	// 예제1

	multiply := func(x, y int) int { //익명함수
		return x * y
	}
	r1 := multiply(10, 5)

	fmt.Println(r1)

	m, n := 5, 10            // 변수가 캡쳐
	sum := func(c int) int { // 익명함수 변수 할당
		return m + n + c // 지역변수 소멸되지 않음. 호출할때마다 생성됨
	}
	r2 := sum(10)

	fmt.Println(r2)
	m = 200
	r2 = sum(10)
	fmt.Println(r2)
}
