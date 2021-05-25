package main

import "fmt"

func main() {
	// iota를 쓴 상수는 밑줄로 생략이 가능함
	const (
		_ = iota
		A
		_
		C
		_
		B
	)

	fmt.Println(A, B, C)
}
