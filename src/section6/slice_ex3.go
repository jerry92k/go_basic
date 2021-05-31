package main

import "fmt"

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make로 공간을 할당 후 복사 해야한다.
	// 복사 된 슬라이스 값 변경해도 원본에는 영향 없음

	// 예제1(복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice2)

	fmt.Println("ex1 :", slice2)
	fmt.Println("ex1 :", slice3)

	fmt.Println()

	// 예제2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 :", a)
	fmt.Println("ex2 :", b)

	//예제 3
	c := [5]int{1, 23, 4, 5, 5}
	d := c[0:3] // 주의! 부분적으로 슬라이스 추출은 참조 -> 원본값 변경된다.
	d[1] = 7

	fmt.Println("ex3 :", c)
	fmt.Println("ex3 :", d)

	//예제 4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 0~5까지 복사하고 용량은 7로 설정해라.

	fmt.Println("ex4 :", e)
	fmt.Println("ex4 :", f)

}
