package main

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) { // 매서드 리턴 값 error 타입 중요
	if n != 0 {
		s := fmt.Sprint("helloGolang :", n)
		return s, nil
	}
	return "", fmt.Errorf("%d를 입력했습니다.", n)
}

func main() {
	// Errorf를 이용한 에러 처리

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", a)

	b, err := notZero(0) // 0 넣고 테스트. 짧은변수 선언으로 동일한 변수명 재 선언해도 됨
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex2 :", b)
	fmt.Println("ex2 end")
}
