package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// RWMutext : 읽기&쓰기 lock : 쓰기 시도 중 다른 곳에서 이전 값 읽으면 안되기 때문에, 다른 스레드에서 읽기, 쓰기 하는 것 모두 방지
	// RMutext : 읽기 lock : 읽기 시도 중에 값 변경 방지. 쓰기 방지만 방지
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex) // var mutext = enw(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 뮤텍스 잠금
			mutex.Lock()
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(200 * time.Millisecond)

			// 쓰기 뮤텍스 잠금 해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {

			//읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
			// 읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
