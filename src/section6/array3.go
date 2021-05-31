package main

import "fmt"

func main() {

	//값 복사 확인 중요

	//예제1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Println(arr1, &arr1)
	fmt.Println(arr2, &arr2)

	//값 자체가 복사되어 주소가 달라지는 것을 확인할 수 있음
	fmt.Printf(" ex1 : %p %v\n", &arr1, arr1)
	fmt.Printf(" ex1 : %p %v\n", &arr2, arr2)
}
