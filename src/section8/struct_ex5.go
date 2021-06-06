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

// 매소드 오버라이딩
func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee
	specialBonus float64
}

func main() {
	//구조체 임베디드 매소드 오버라이딩 패턴

	ep1 := Employee{"kim", 2000, 51504}
	ep2 := Employee{"park", 12323, 4444}

	ex := Executives{
		Employee{"lee", 23434, 3333}, // is a 관계
		20000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// 구조체를 통해서 매소드 호출
	fmt.Println("ex2 : ", int(ex.Calculate()))

	// 오버라이딩한 매소드를 사용하지 않고 Employee 매소드 사용
	fmt.Println("ex2 : ", int(ex.Employee.Calculate()+ex.specialBonus))

}
