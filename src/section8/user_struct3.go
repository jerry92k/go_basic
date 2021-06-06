package main

import "fmt"

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt : %d, price %d, orderprice :%d", cnt, price, fn(cnt, price))
}

func main() {
	var orderPrice totCost
	orderPrice = func(cnt int, price int) int {
		return (cnt * price) * 10000
	}

	describe(3, 5000, orderPrice)
}
