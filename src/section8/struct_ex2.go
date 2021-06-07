package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {

	kim := Account{"245-902", 10000000, 0.53}
	lee := Account{"245-901", 20000000, 0.53}

	fmt.Println(kim)
	fmt.Println(lee)

	CalculateD(kim)
	//receiver인 경우 자동으로 참조형으로 변환되지만 함수인 경우는 참조타입을 넘겨줘야함
	CalculateP(&lee)

	fmt.Println(kim.balance)
	fmt.Println(lee.balance)

}
