// duck typing

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

type Bird struct {
	name   string
	weight int
}

// 구조체 Dog 매소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, "Dog bites!")
}

func (d Dog) sounds() {
	fmt.Println(d.name, "Dog barks!")
}

func (d Dog) run() {
	fmt.Println(d.name, "Dog runs!")
}

func (d Cat) bite() {
	fmt.Println(d.name, "Cat bites!")
}

func (d Cat) sounds() {
	fmt.Println(d.name, "Cat barks!")
}

func (d Cat) run() {
	fmt.Println(d.name, "Cat runs!")
}

func (d Bird) bite() {
	fmt.Println(d.name, "Bird bites!")
}

func (d Bird) sounds() {
	fmt.Println(d.name, "Bird barks!")
}

// 동물의 행동 인터페이스 선언. 매서드 만으로 그 객체를 판단 => duck typing
type Behaivor interface {
	bite()
	sounds()
	run()
}

func act(animal Behaivor) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func main() {
	// 예제1
	dog := Dog{"ppoya", 10}
	cat := Cat{"poll", 2}
	bird := Bird{"ggg", 4}

	act(dog)
	act(cat)
	act(bird)

}
