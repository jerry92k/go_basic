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
	CalculateP(&lee)

	fmt.Println(kim.balance)
	fmt.Println(lee.balance)

}
