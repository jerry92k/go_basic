package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 객체
	// 동기화 상태(조건) 매소드 사용
	// 다른 언어 : [wait, notify, notifyall]
	// 고루틴 : wait, signal, broadcast

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine wating", n)
			condition.Wait()
			fmt.Println("waiting end", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		//	<-c
		fmt.Println("received :", <-c)

	}
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("wake goroutine(signal) :", i)
		condition.Signal()
		mutex.Unlock()
	}

	time.Sleep(2 * time.Second)

	mutex.Lock()
	fmt.Println("wake goroutine (broadcast)")
	condition.Broadcast()
	mutex.Unlock()
	time.Sleep(2 * time.Second)

}
