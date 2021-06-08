package main

import (
	"errors"
	"fmt"
)

func main() {
	// errors 패키지의 New 매서드를 활용한 에러 생성
	// Errorf 보다 더 많이 사용

	var err1 error = errors.New("Error occoured -1")
	err2 := errors.New("Error occured -2")

	fmt.Println("error1 : ", err1)
	fmt.Println("error1 : ", err1.Error())
	fmt.Println("error1 : ", err2)
	fmt.Println("error1 : ", err2.Error())
}
