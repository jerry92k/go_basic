package main

import "fmt"

func main() {
	//포인터
	//Go : 포인터 지원(c)
	// 사용 이유 : 변수의 지역성, 연속된 메모리의 참조, 힙, 스택...
	// 파이썬, 자바, C# 등은 포인터 지원하지 않음. but 자바 JRE-> 컴파일러, 인터프리터 등에서 사용
	//  => 가비지 컬렉터로 jvm 같은 머신에서 처리
	// Go에서 주소의 값을 직접 변경 불가능 -> 잘못된 코딩으로 인한 버그 방지
	// *(에스터리스크)로 사용
	// nil로 초기화 (nil==0)

	//예제 1. 선언방법
	var a *int //방법1.
	var b *int = new(int)

	fmt.Println(a) // nill로 초기화 상태
	fmt.Println(b) // 메모리 할당된 상태

	i := 7

	fmt.Println("ex1 : ", i) // i값
	fmt.Println(&i)          // i의 주소

	a = &i
	b = &i

	fmt.Println("ex1-2 : ", a, &i)
	fmt.Println("ex1-2 : ", &a)
	fmt.Println("ex1-2 : ", *a) // 역참조

	fmt.Println("ex1-3 : ", b, &i)
	fmt.Println("ex1-3 : ", &b)
	fmt.Println("ex1-3 : ", *b)

	var c = &i
	d := &i

	*d = 10

	fmt.Println("ex1-2 : ", c, &i)
	fmt.Println("ex1-2 : ", &c)
	fmt.Println("ex1-2 : ", *c) // 역참조

	fmt.Println("ex1-3 : ", d, &i)
	fmt.Println("ex1-3 : ", &d)
	fmt.Println("ex1-3 : ", *d)

}
