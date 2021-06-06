package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

// bite 매소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

// 동물의 행동 인터페이스 선언
type Behaivor interface {
	bite()
}

func main() {
	dog1 := Dog{"poll", 10}

	var interface1 Behaivor

	interface1 = dog1
	interface1.bite()

	dog1.bite()

	// 예제 2
	dog2 := Dog{"marry", 12}
	interface2 := Behaivor(dog2)

	interface2.bite()

	interface3 := []Behaivor{dog1, dog2}

	// 인덱스 형태로 실행
	for idx, _ := range interface3 {
		interface3[idx].bite()
	}

	// 값 형태로 실행(인터페이스)

	for _, val := range interface3 {
		val.bite()
	}
}
