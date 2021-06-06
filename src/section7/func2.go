package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a int, b int) {
	fmt.Println("ex1 :", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i = *i + 100
}

func main() {
	//함수(콜백) 전달)

	//예제1
	sum(100, add)

	//예제2
	a := 100
	multi_value(a)
	fmt.Println(a)
	multi_reference(&a)
	fmt.Println(a)
}
