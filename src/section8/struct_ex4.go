package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee
	specialBonus float64
}

func main() {
	//구조체 임베디드 타입
	// 매소드를 대 사용하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 매소드 상속을 활용한 패턴

	ep1 := Employee{"kim", 2000, 51504}
	ep2 := Employee{"park", 12323, 4444}

	ex := Executives{
		Employee{"lee", 23434, 3333},
		20000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// 구조체를 통해서 매소드 호출
	fmt.Println("ex2 : ", int(ex.Calculate()+ex.specialBonus))

}
