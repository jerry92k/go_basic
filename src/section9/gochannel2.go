package main

import "fmt"

func rangeSum(rg int, c chan int) {
	sum := 0

	for i := 0; i <= rg; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	// 채널
	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(700, c)
	go rangeSum(200, c)

	// 작업이 끝나는 순서대로 데이터 수신
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("ex 1 :", result1)
	fmt.Println("ex 1 :", result2)
	fmt.Println("ex 1 :", result3)

}
