// 파일쓰기
package main

import (
	"encoding/csv"
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
	// csv 파일 쓰기 예제
	// 패키지 저장소를 통해서 excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Create("test_write.csv")
	errCheck1(err)

	//리소스 해제
	defer file.Close()

	// csv writer 생성
	wr := csv.NewWriter(file)
	// wr:=csv.NewWriter(bufio.NewWriter(file))

	wr.Write([]string{"kim", "4.8"})
	wr.Write([]string{"Lee", "4.2"})
	wr.Write([]string{"Park", "4.3"})
	wr.Write([]string{"Cho", "4.9"})
	wr.Write([]string{"Hong", "4.2"})

	wr.Flush() // 버퍼 -> 파일로 쓰기

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기 (%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명", fi.Name())
	fmt.Println("CSV 파일 권한", fi.Mode())
}
