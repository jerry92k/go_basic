package main

import "fmt"

func main() {
	cnt := increaseCnt()

	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
}

func increaseCnt() func() int {
	n := 0 // 지역번수(캡쳐됨
	return func() int {
		n += 1
		return n
	}
}
