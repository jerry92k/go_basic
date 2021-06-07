package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 채널
	// 예제1 (비동기 : 버퍼 사용)
	// 버퍼 : 발신 -> 가득차면 대기, 비어있으면 작동, 수신 -> 비어있으면 대기, 가득차면 작동

	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 3) // channel의 용량을 2로 설정 // 2개 보내고 2개 받기
	cnt := 12

	// go routine 스레드에서 ch 로 데이터를 보내고
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("go : ", i)
		}
	}()

	// main thread 에서 ch에 수신된 데이터를 출력
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("main :", i)
	}
}
