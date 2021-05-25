package main

import "fmt"

func main() {

	//예제1(배열)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)

	fmt.Printf(" ex1 : %p %v\n", &arr1, arr1)
	fmt.Printf(" ex1 : %p %v\n", &arr2, arr2)

	//예제2 : 슬라이스
	slice1 := []int{1, 2, 3}
	var slice2 []int
	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex1 :", slice1)
	fmt.Println("ex2 :", slice2)

	// 예제3. 슬라이스 예외상황

	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4])
	//	fmt.Println("ex3 : ", slice3[50]) // 용량만큼 초기화 되는 것이 아니고 길이만큼 초기화됨
}
