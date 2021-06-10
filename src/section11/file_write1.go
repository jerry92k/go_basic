// 파일쓰기
package main

import (
	"fmt"
	"os"
)

//에러 체크 방식1

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

// 에러체크방식2
func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	// 파일 쓰기
	// Create: 새 파일 작성 및 파일 열기
	// Close: 리소스 닫기
	// Write, WriteString ,WriteAt : 파일 쓰기
	// 각 운영체제 권한 주의(오류메세지 확인)
	// 예외 처리 정말 중요!
	// https://golang.org/pkg/os

	// 파일 쓰기 예제
	file, err := os.Create("test_write.txt")
	errCheck1(err)

	defer file.Close()

	// 쓰기 예제1
	s1 := []byte{48, 49, 50, 51, 52}
	n1, err := file.Write(s1)

	errCheck2(err)

	fmt.Println("쓰기 작업 완료", n1)

	file.Sync() // Write Commit

	// 쓰기 예제2
	s2 := "Hello Golang \n File write test! -1 \n"
	n2, err := file.WriteString(s2)
	errCheck2(err)

	fmt.Println("쓰기 작업 완료- 2", n2)

	file.Sync()

	s3 := "Test WriteAt -2\n"
	n3, err := file.WriteAt([]byte(s3), 70) // len(offset) 조절

	errCheck1(err)

	fmt.Println("쓰기 작업 완료- 3", n3)

	// 쓰기 예제 4
	n4, err := file.WriteString("Hello Golang! \n File Write test -3 \n")
	errCheck1(err)

	fmt.Println("쓰기 작업 완료- 4", n4)
}
