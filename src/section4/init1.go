package main

import (
	"fmt"
)

func init() {
	// init : 패키지 로드시에 가장 먼저 호출
	// 가장 먼저 초기화되는 작업 작성 시 유용하다.
	fmt.Println("init")
}

func main() {

	fmt.Println("main method start")
}
