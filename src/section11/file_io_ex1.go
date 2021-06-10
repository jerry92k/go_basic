package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 활용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	// 파일 열고 닫고 별도로 안해도 됨
	// 아래 매소드 확인 가능
	// WriteFile(), ReadFile(), ReadAll() 등 활용
	// https://golang.org/pkg/io/ioutil

	s := "hello golang \n file write test! \n"

	// 파일 모드 -> 퍼미션
	// 읽기 : 4, 쓰기 : 2, 실행 :1
	// 소유자|그룹|기타사용자 ( ex.644 -> 소유자는 읽기, 쓰기 가능. 그룹과 기타사용자는 읽기만 가능)
	// https://golang.org/pkg/os/#FileMode

	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644)) // go에서 퍼미션을 줄때는 앞에 0을 넣음
	errCheck(err)

	data, err := ioutil.ReadFile("sample.txt")
	fmt.Println("========================")
	fmt.Println(data)
	fmt.Println(string(data))
	fmt.Println("========================")
}
