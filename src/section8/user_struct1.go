package main

import "fmt"

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

//구조체 <->매소드 바인딩
// 일반메소드
func price(c Car) int64 {
	return c.price + c.tax
}

func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {
	// Go -> 객체지향타입을 구조체로 정의한다. ( 클래스, 상속 개념 없음)
	// 객체 지향 -> 클래스(속성, 기능(상태:메소드)) : 코드의 재사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그래밍
	// 객체지향 언어일까? 맞다!
	// Go 는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스를 통한 다형성 지원, 구조체를 클래스형태의 코딩 가능
	// 객체지향의 기본 개념 -> Go에서 포함하고 있다 -> 객체지향언어
	// 상태, 메소드 분리해서 정의 (결합성 없음)
	// 사용자 정의 타입 : 구조체, 인터페이스, 기본타입(int, float, string...) 함수
	// 구조체 메소드 연결을 통해서 타언어의 클래스형식처럼 사용 가능(객체지향)

	// 예제1
	bmw := Car{name: "520d", price: 500000, color: "white", tax: 50000}
	bmw2 := Car{name: "720d", price: 500000, color: "white", tax: 50000}
	benz := Car{name: "220d", price: 1000000, color: "blue", tax: 10000}

	fmt.Println("ex 1 ", bmw, &bmw)
	fmt.Println("ex 1 ", benz, &benz)

	fmt.Printf("ex 1 : %p", &bmw)
	fmt.Printf("ex 1 : %p", &benz)

	fmt.Println("ex2 : ", price(bmw))
	fmt.Println("ex2 : ", bmw.Price())

	// 동일 구조체로 생성된 객체의 주소 확인
	fmt.Printf("ex3 : %p", &bmw)
	fmt.Printf("ex3 : %p", &bmw2)
}
