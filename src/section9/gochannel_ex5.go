package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	수신
	go func() {
		for {
			num := <-ch1
			fmt.Println("ch1 :", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	// 발신
	go func() {
		for {
			ch2 <- "hi go~"
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case ch1 <- 777: // 발신용도
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
