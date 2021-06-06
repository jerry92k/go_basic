package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 타입 변환 -> Type Assertion
	// 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변경 후 사용해야 하는 경우
	// 인터페이스. (타입) -> 형 변환
	// interfaceVal. (type)

	// 예제1
	var a interface{} = 15
	b := a
	c := a.(int)
	//	d := a.(float64)

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
	//	fmt.Println(d) // 런타임시 타입 변환 오류남

	// 예제2(자장된 실제 타입검사)
	if v, ok := a.(int); ok {
		fmt.Println(v, ok)
	}
}
