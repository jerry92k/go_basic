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

	//인터페이스 규격화 역할 이해
	// 인터페이스에 정의된 매소드 사용 유도
	// 코드의 가독성 및 유지보수 증가

	//덕타이핑 예제
	// 덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현 매서드로만 판단하는 방식
	// Go의 중요한 특징 : 오리처럼 걷고, 소리내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다 -> 덕타이 의미

	// 예제1
	dog := Dog{"ppoya", 10}
	cat := Cat{"poll", 2}
	// bird := Bird{"ggg", 4} // Behaivor의 모든 매서드를 구현하지 않았으므로 오류

	act(dog)
	act(cat)
	act(bird)

}
