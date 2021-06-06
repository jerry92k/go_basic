package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	// 구조체 : 필드는 갖지만 매소드는 갖지 않는다.
	// 객체지향 방식을 지원 -> 리시버를 통해 매소드랑 연결
	// 상속, 객체, 클래스 개념 없음
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	// 예제1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}
	lee := Account{number: "245-902", balance: 20000000}
	park := Account{number: "245-903", interest: 0.025}
	cho := Account{"245-904", 15000000, 0.4}

	fmt.Println(kim)
	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(cho)

	// 예제2
	fmt.Println(int(kim.Calculate()))
	fmt.Println(int(lee.Calculate()))
	fmt.Println(int(park.Calculate()))
	fmt.Println(int(cho.Calculate()))
}
