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
	// 다양한 선언 방법
	// &struct, struct : &struct 포인터를 받아오기 위해 역참조를 하기 때문에 속도는 조금 느리다
	// 인터페이스 매소드를 선언만 해둔 후 -> 오버라이딩 해서 매서드에 포인터 리시버를 사용할 경우 반드시 &struct

	// 예제1
	// 선언 방법1
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 1000000000
	kim.interest = 0.34

	// 선언 방법2
	hong := &Account{number: "245-902", balance: 15000000, interest: 0.213}

	// 선언 방법3
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 213434233
	lee.interest = 0.213434233

	fmt.Println(kim)
	fmt.Println(hong)
	fmt.Println(lee)

	fmt.Printf("ex : %#v\n", kim)
	fmt.Printf("ex : %#v\n", hong)
	fmt.Printf("ex : %#v\n", lee)

	//예제2
	fmt.Println("ex 2 : ", int(kim.Calculate()))
	fmt.Println("ex 2 : ", int(hong.Calculate()))
	fmt.Println("ex 2 : ", int(lee.Calculate()))

}
