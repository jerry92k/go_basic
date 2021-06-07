package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널
	// 예제1 (동기 : 버퍼 미사용)

	ch := make(chan bool)
	cnt := 6

	// go routine 스레드에서 ch 로 데이터를 보내고
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("go : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	// main thread 에서 ch에 수신된 데이터를 출력
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("main :", i)
	}
}
