package main

import "fmt"

//포인터 값 전달 => 원본 값 변경 위해서 사용
// 특히 크기가 큰 배열인 경우 값 복사시에 시스템 부담을 해결 (이 문제에서 슬라이스, 맵은 참조로 전달되기 때문에 상관없음 )
func rptc(n *int) {
	*n = 77
}

// 함수, 매서드 호출 시에 매개변수 값을 복사 전달 => 원본 값 변경 불가능
func vptc(n int) {

}

func main() {

	var a int = 10
	var b int = 10

	fmt.Println("ex1 :", a)
	fmt.Println("ex1 :", b)

	rptc(&a)
	vptc(b)

	fmt.Println("ex1-2 :", a)
	fmt.Println("ex1-2 :", b)

}
