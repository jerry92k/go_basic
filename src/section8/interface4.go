package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func (d Dog) run() {
	fmt.Println(d.name, "Dog runs!")
}

func (d Cat) run() {
	fmt.Println(d.name, "Cat runs!")
}

func act(animal interface{ run() }) { // 익명 인터페이스. 타입 정의 X
	animal.run()
}

func main() {

	// 예제1
	dog := Dog{"ppoya", 10}
	cat := Cat{"poll", 2}
	// bird := Bird{"ggg", 4} // Behaivor의 모든 매서드를 구현하지 않았으므로 오류

	act(dog)
	act(cat)

}
