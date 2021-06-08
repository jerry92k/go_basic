package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 에러 처리
	// 소프트웨어 품질 향상 가장 중 -> 유형코드 및 에러 정보 등의 정보를 남기는 것
	// Go에서는 기본적으로 error 패키지에서 error 인터페이스 제공
	// type error interface{ Error() string}
	// 즉, Error 메서드를 구현하면 사용자 정의 에러 처리 제작 가능
	// 기본적으로 매서드마다 리턴 타입 2개(리턴값, 에러)
	// 주로 Errorf(에러 타입 리턴) 매서드, fatal(프로그램 종료) 매서드를 통해서 에러 출력

	// 기본적인 매서드 에러 처리 예제
	f, err := os.Open("unnamedfile") // 파일 열기 매서드 -> 없는 이름 에러
	if err != nil {                  // nil 이면 에러가 없음
		log.Fatal(err.Error()) // 방법 1
		// log.Fatal(err) // 방법 2. 방법1과 결과 똑같음
	}

	fmt.Println("================")
	fmt.Println("ex 1 :", f.Name())
}
