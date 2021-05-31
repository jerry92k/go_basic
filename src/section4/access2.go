package main

import (
	"fmt"
	checkUp "section4/lib"
	checkUp2 "section4/lib2"
)

func main() {

	fmt.Println("100보다 큰수? :", checkUp.CheckNum(101))

	fmt.Println("1000보다 큰수? :", checkUp2.CheckNum2(200))
	fmt.Println(checkUp2.CheckNum3())
}
