package main

import "fmt"

func main() {
	// IF문 : 반드시 Boolean 검사. 1,0 사용불가 ( 자동 형변환 불가)
	// 소괄호 사용 // X

	var a int = 20
	b := 20

	//예제 1
	if a >= 15 {
		fmt.Println("25이상")
	}

	// 예제 2
	//if b>=25
	if b >= 25 {
		fmt.Println("괄호를 붙이지 않으면 if문 바로 뒤세 세미콜론을 붙여버려 syntax 에러 남")
	}

	if c := true; c {
		fmt.Println("True")
	}

	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}
}
