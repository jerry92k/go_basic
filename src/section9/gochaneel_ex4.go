package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널 셀럭트 구분
	// 채널 select 구문 -> 채널에 값이 수신되면 해당 case 문 쉴행
	// 일회성 구문이므로, for(반복문) 안에서 수행
	// default 구문 처리 주의

	// 예제1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "go hi~"
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
