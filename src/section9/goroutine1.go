// 고루틴 기초
package main

import (
	"fmt"
	"time"
)

func ex1() {
	fmt.Println("exe1 func start", time.Now())
	//time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func ex2() {
	fmt.Println("exe2 func start", time.Now())
	//time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}

func ex3() {
	fmt.Println("exe3 func start", time.Now())
	//	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}

func main() {
	// 고루틴(Goroutine)
	// 타 언어의 스레드(thread)와 비슷한 기능
	// 생성 방법 매우 간단. 리소스 매우 적게 사용. 다른 언어는 MB 단위로 리소스를 할당받지만 go는 KB 단위로 할당받음
	// 즉, 수많은 고루틴 동시 생성 실행 가능
	// 비동기적 함수 루틴 실행() -> 채널을 통한 통신 가능
	// 공유메모리 사용 시에 정확한 동기화 코딩 필요
	// 싱글루틴에 비해 항상 빠른 처리 결과는 아니다

	// 멀티스레드 장점과 단점
	// 장점 : 응답성이 향상, 자원공유를 효율적으로 활용 사용, 작업이 분리되어 코드 간결
	// 단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스에 사이드 이펙트 발생 (ex. 성능저하), 동기화 코딩 반드시 숙지 필요

	ex1() // main routine

	fmt.Println("main routine start", time.Now())

	go ex2()                    // sub routine
	go ex3()                    // sub routine
	time.Sleep(4 * time.Second) // main 이 끝나면 subroutine이 끝나기 때문에 sleep을 줌
	fmt.Println("main routine end", time.Now())

}
