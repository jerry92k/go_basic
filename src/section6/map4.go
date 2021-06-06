package main

import "fmt"

func main() {
	//맵 조회할 경우 주의할 점

	map1 := map[string]int{ //데이터 기본값 초기화 : int : 0, string:"", float:0.0
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"] // ok에는 key가 존재하는지 안하는지에 대한 부울값이 리턴됨.

	fmt.Println("ex1 :", value1)
	fmt.Println("ex1-2 :", value2)
	fmt.Println("ex1-3 :", value3)
	fmt.Println("ex1-4 :", ok)

	//예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exist")
	}
}
