package main

import (
	"fmt"
)

func main() {
	// 예제1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println(" a는 짝수")
	case 1, 3, 5:
		fmt.Println(" a는 홀수")
	}

	// 예제3.
	// fallthrough  => 조건에 맞는 문장에 해당 예약어가 있다면 그 이후에는 조건에 맞지 않더라도 실행됨
	switch e := "go"; e {
	case "java":
		fmt.Println("java다")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
	}
}
