package main

import (
	"fmt"
	"reflect"
)

// 앞글자를 대문자로 사용하면 외부에서 참조 가능. 소문자로 사용하면 외부에서 참조 불가
type Car struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func main() {
	// 필드 태그 사용

	// 타입 선언 뒤에 명시한 것

	// 에제 1
	tag := reflect.TypeOf(Car{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("ex1 :", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}
