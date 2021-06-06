package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a *Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {

	// 정리 : 구조체 인스턴스 값 변경 시 -> 포인터 전달. 보통의 경우-> 값 전.
	kim := Account{"245-902", 10000000, 0.1}

	fmt.Println(int(kim.balance))

	CalculateD(kim)
	//receiver인 경우 자동으로 참조형으로 변환되지만 함수인 경우는 참조타입을 넘겨줘야함

	fmt.Println(int(kim.balance))

	kim.CalculateD(9999999)
	fmt.Println(int(kim.balance))

}
