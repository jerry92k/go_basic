package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}
func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}
	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println(value)
	}
}

func main() {
	// 가변인자 실습(매개변수 개수가 동적으로 변할때 - 정해져있지 않음)
	x := multiply(1, 2, 3)
	y := sum(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(x)
	fmt.Println(y)

	prtWord("a", "apple", "gogogo")

	// 슬라이스 인자 넘기기
	a := []int{5, 6, 6, 7, 8, 9, 10}
	m := multiply(a...)
	fmt.Println(m)
	fmt.Println(sum(a...))
}
