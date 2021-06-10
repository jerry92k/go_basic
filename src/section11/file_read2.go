// 파일쓰기
package main

import (
	"bufio"
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
	// 파일 읽기
	// csv 파일 읽기 예제

	file, err := os.Open("sample.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	//rr :=csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) // 권장

	row1, err1 := rr.Read() // row 한줄 읽어옴
	errCheck1(err1)
	row2, err2 := rr.Read()
	errCheck2(err2)
	fmt.Println("CSV Read Example")
	fmt.Println(row1)
	fmt.Println(row1[0], row1[1], row1[2], row1[3], row1[1:5])
	fmt.Println(row2[0], row2[1], row2[2], row2[3], row2[1:5])

	fmt.Println("================================")

	rows, err := rr.ReadAll() // 전체 Row 읽기
	errCheck2(err)
	fmt.Println("CSV ReadAll Example")
	fmt.Println(rows)
	fmt.Println(rows[5][1], rows[2])
	fmt.Println("================================")

	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s  ", rows[i][j])
		}
		fmt.Println()
	}
}
