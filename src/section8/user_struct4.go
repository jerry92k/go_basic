package main

import "fmt"

type shoppingBasket struct{ cnt, price int }

// 결제함수
func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본수정(참조형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본수정(값전달형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// 리시버(구조체 매소드) 전달(값,참조)형식
	// 함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달
	// 리시버(구조체)도 마찬가지로 포인터를 활용해서 매소드 내에서 원본 수정 가능

	bs1 := shoppingBasket{3, 5000}
	fmt.Println("ex1 :", bs1.purchase())

	bs1.rePurchaseD(7, 5000)
	fmt.Println("ex1-1 :", bs1.purchase())

	bs1.rePurchaseP(3, 15000)
	fmt.Println("ex1-2 :", bs1.purchase())
}
