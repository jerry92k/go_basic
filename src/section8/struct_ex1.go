package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest} // 구조체 인스턴스를 생성한 뒤 포인터를 리턴
}

func main() {
	// 구조체 생성자 패턴 예제

	// 예제1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.0125}

	var lee *Account = new(Account)
	lee.number = "245-902"
	lee.balance = 1232302
	lee.interest = 20.2321

	fmt.Println("ex 1 :", kim)

	// 예제2 : 생성자를 사용하는것과 같은 느낌
	park := NewAccount("245-903", 1272321, 0.04)
	fmt.Println("ex 1 :", park)
	fmt.Printf("ex 1 :%#v", park)
	fmt.Printf("ex 1 :%#v", kim)

}
