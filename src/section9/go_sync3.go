package main

import (
	"fmt"
	"time"
)

func main() {
	// 뮤텍스 : 상호 배제 -> Thread(고루틴) 들이 런타임에 서로 영향 주지 않게 보장해주는 기술
	// 뮤텍스 : 여러 고루틴에서 작업하는 공유데이터 보호

	// 동기화 사용하지 않은 경우 에제
	// 쓰기 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환할 가능성 증가
	// 시스템 전체 cpu 사용

	data := 0
	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
