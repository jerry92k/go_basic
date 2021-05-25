package main

import (
	"fmt"
	"section4/lib"
)

var num int32

func init() {
	num = 30
	fmt.Println("init")
}

func main() {
	fmt.Println("10보다 큰 수 ? : ", lib.CheckNum(num))
}
