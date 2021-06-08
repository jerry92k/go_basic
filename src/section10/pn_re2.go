// panic & recover
package main

import (
	"fmt"
)

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error Message :", s)
	}()

	panic("Error occurred! ")

	fmt.Println("여기는 호출이 안됨") //panic이 발생하며 실행은 중단되고, defer 함수 호출 후 자기자신을 호출한 main으로 리턴되므로
}

func main() {
	// panic으로 발생한 에러를 복구 후 코드 재 실행(프로그램 종료되지 않는다.)
	// 즉, 코드 흐름을 정상상태로 복구하는 기능
	// panic에서 설정한 메시지를 받아올 수 있다.

	//예제
	runFunc()

	fmt.Println("Hello golang")

}
