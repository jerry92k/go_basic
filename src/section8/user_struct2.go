package main

import "fmt"

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입
	// 예제1

	a := cnt(5)

	//예제2

	var b cnt = 15

	fmt.Println(a)
	fmt.Println(b)

	//	testConverT(a) // 타입 오류!
	testConverD(a)
	testConverT(int(a))

}

func testConverT(i int) {
	fmt.Println("ex3 : (Default Type) :", i)
}

func testConverD(i cnt) {
	fmt.Println("ex3 : (Default Type) :", i)
}
