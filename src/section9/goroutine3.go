package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start : ", time.Now())

	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())

}

func main() {
	// 고루틴
	// 멀티코어(다중 cpu) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU())                       //런타임에 현재 사용하는 CPU 코어
	fmt.Println("current system CPU :", runtime.GOMAXPROCS(0)) // 설정값 출력

	//예제1
	fmt.Println("main routine start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("main routine end : ", time.Now())
}
