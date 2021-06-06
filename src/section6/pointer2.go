package main

import "fmt"

func main() {

	//예제 1
	i := 7
	var p = &i

	fmt.Println("ex1 :", i, *p, &i, p)

	*p++ // 주소값을 증가시키는게 아니고 역참조하여 값을 증가하는 거여서 가능

	fmt.Println("ex1 :", i, *p, &i, p)

	i = 77

	fmt.Println("ex1 :", i, *p, &i, p)
}
