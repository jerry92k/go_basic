//함수 기초

package main

import "fmt"
import "strconv"

// 함수 선언 위치는 어느 곳이든 가능
func helloGolang() {
	fmt.Println("hello golang")
}

func say_one(m string) {
	fmt.Println("ex : ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {

	helloGolang()
	say_one("say one")
	result := sum(5, 5)
	fmt.Println(result)
	fmt.Println(sum(2, 5))
	fmt.Println("ex :", strconv.Itoa(sum(5, 5)))
	// 함수
	// 선언 : func 키워드로 선언
	// func 함수명(매개변수) (반환타입 or 반환값 변수명) : 매개변수 존재, 반환값 존재
	// func 함수명() (반환타입 or 반환값 변수명) : 매개변수 없음, 반환값 존재
	// func 함수명(매개변수) :  매개변수 존재, 반환값 없음
	// func 함수명() :  매개변수 없음, 반환값 없음
	// 타 언어와 달리 반환값(return value) 다수 가능

	// 예제1

}
