// 채널 기초 1

package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("work1 : s --> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work1 : E --> ", time.Now())

	v <- 1
}

func work2(v chan int) {
	fmt.Println("work2 : s --> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work2 : E --> ", time.Now())

	v <- 2
}

func main() {
	// 채널은 참조형(reference type)
	// 고루틴 간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용
	// 실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
	// 데이터 전달 자료형 선언 후 사용(지정된 타입만 주고받을 수 있음)

	// interface{} 전달을 통해서 자료형 상관없이 전송 및 수신 가능
	// 값을 전달(복사 후 : bool, int등), 포인터(슬라이스, 맵) 등을 전달시에는 주의! > 동기화 사용(Mutex)
	// 멀티프로세싱 처리에서 교착상태(경쟁) 주의!

	// <-, -> (채널 <- 데이터 : 송신), (변수 <- 채널 : 수신)

	// 예제 1

	fmt.Println("main : s--> ", time.Now())

	// var c chan int
	// c = make(chan int)

	v := make(chan int) // int형 채널 선언. 채널은 참조형
	go work1(v)
	go work2(v)

	<-v
	<-v

	// 채널은 동기식이라 Sleep을 하지 않아도 기다림

	fmt.Println("main : E--> ", time.Now())

}
