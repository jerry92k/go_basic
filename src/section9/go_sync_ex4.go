package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	// 원자성 사용 -> 기능적으로 불가능한 완전 보증된 일련의 프로세스.
	// 모두 성공 or 모두 실패
	// 모든 조작이 완료될 때가지 다른 프로세스 개입 불가
	// sync/atomic 에서 원자적 연산자 제공
	// https://golang.org/pkg/sync/atomic 에서 계열 확인 가능
	// 주로 공용 변수에 관한 계산 사용

	// 원자성 사용 안할 경우 예제
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0

	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			//fmt.Println("wait Group : ", n)
			//			cnt += 1
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			//fmt.Println("wait Group : ", n)
			//		cnt -= 1
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수 같아야 함

	wg.Wait()

	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("WaitGroup end!", cnt)
	fmt.Println("WaitGroup end!", finalCnt)
}
