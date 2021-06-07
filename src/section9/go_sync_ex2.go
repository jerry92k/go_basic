package main

import (
	"fmt"
	"sync"
)

func main() {

	// 대기 그룹
	// 실행 흐름 변경 (고루틴이 종료될 때 까지 대기 기능)
	// waitGroup : Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시 까지 대기)
	// Add로 추가된 고루틴 개수와 Done으로 종료되는 알림 횟수 같아야 정확하게 동작 (Add == Done)

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("wait Group : ", n)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수 같아야 함

	wg.Wait()
	fmt.Println("WaitGroup end!")
}
