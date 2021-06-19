package main

import "fmt"

var closurFunc func(int) int

func outerFunc() {
	m, n := 5, 10
	closurFunc = func(c int) int { // 익명함수 변수 할당
		return m + n + c // outerFunc가 종료되어 메모리가 반환되어도 m, n 지역변수에 대한 참조가 살아있어 GC에서 메모리 해제가 되지 않음
	}
	r2 := closurFunc(10)
	fmt.Println("here ", r2)
}

func main() {
	outerFunc()
	r2 := closurFunc(20)
	fmt.Println("here ", r2)
}
