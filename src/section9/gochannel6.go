package main

import (
	"fmt"
)

func main() {
	// 채널(channel)
	// Close : 채널 닫기

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- "good!"
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex1 : ", val3, ok3)

	close(ch)

	val4, ok4 := <-ch
	fmt.Println("ex1 : ", val4, ok4) // 채널이 닫혀서 값을 수신하지않았다. 알 수 있음. 이것을 통해 채널이 닫혔나 확인 가능
}
