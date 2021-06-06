package main

import "fmt"

type Car struct { // 첫 글자 대문자로 선언
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct { // 첫 글자 소문자로 선언
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	// 중첩 구조체

	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}

	fmt.Println("ex1 :", car1.name)
	fmt.Println("ex1 :", car1.color)
	fmt.Println("ex1 :", car1.company)
	fmt.Printf("ex1 : %#v \n", car1.detail)

	fmt.Println("ex1 :", car1.detail.length)
	fmt.Println("ex1 :", car1.detail.height)
	fmt.Println("ex1 :", car1.detail.width)
}
