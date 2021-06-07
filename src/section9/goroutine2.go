package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println(name, "start : ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, " >>>> ", i)
	}
	fmt.Println(name, " end ", time.Now())
}

func main() {
	exe("t1")

	fmt.Println("main routine start :", time.Now())

	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second)
	fmt.Println("main routine end :", time.Now())
}
