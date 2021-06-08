package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외(에러) 처리
type PowError struct {
	time    time.Time // 에러 발생 시간
	value   float64   // 파라미터
	message string    // 에러 메세지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - input value (value : %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다"}
	}

	if math.IsNaN(i) { // 숫자가 아니라면
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다"}
	}

	return math.Pow(f, i), nil
}

func main() {
	// error 타입이 아닌 경우 에러 처리 방법
	// Error 매서드를 구현해서 사용자 정의 에러 처리 예제 심화
	// 구조체를 사용해서 세부적인 정보 출력

	// 예제 1
	v, err := Power(10, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex 1: ", v)

	// 예제 2
	t, err := Power(0, 3)
	if err != nil {
		log.Fatal(err.Error())
		fmt.Println(err.(PowError).value)
		fmt.Println(err.(PowError).message)
		fmt.Println(err.(PowError).time)
	}

	fmt.Println("ex 2: ", t)
}
