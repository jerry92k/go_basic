// panic & recover
package main

import (
	"fmt"
)

func runFunc() {
	defer func() { // 지연함수. defer 함수는 함수가 종료되지 전에 실행됨
		if s := recover(); s != nil { // 오류 발생하지 않으면 실행되지 않음
			fmt.Println("Error Message :", s)
		}
	}()

	a := [3]int{1, 2, 3}

	// 패닉을 일부러 일으키지 않고, 런타임시 오류 발생 예시를 만들어봄
	for i := 0; i < 5; i++ {
		fmt.Println("ex 1: ", a[i])
	}
}

func main() {
	// panic으로 발생한 에러를 복구 후 코드 재 실행(프로그램 종료되지 않는다.)
	// 즉, 코드 흐름을 정상상태로 복구하는 기능
	// panic에서 설정한 메시지를 받아올 수 있다.

	//예제
	runFunc()

	fmt.Println("Hello golang")

}
