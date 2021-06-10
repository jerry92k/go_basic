package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	// 외부 저장소 패키지 설치
	// 2가지 방법
	// 1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	// 2. go get 패키지 주소 설치 -> 선언

	// 선언 후 go get 설치 예제 ( 엑셀 파일 읽기 )

	xfile := "sample2.xlsx"

	xlFile, err := xlsx.OpenFile(xfile)

	if err != nil {
		fmt.Println(err)
		panic("Excel load fail")
	}

	for _, sheet := range xlFile.Sheets {
		for rowIdx := 0; rowIdx < sheet.MaxRow; rowIdx++ {
			row, err := sheet.Row(rowIdx)
			if err != nil {
				panic("row load fail")
			}
			//for _, cell := range sheet.Row(rowIdx).MaxCell {
			row.ForEachCell(cellVisitor)

			fmt.Println("")
		}
	}
}

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		text := value
		fmt.Printf("%s\t", text) // 각 셀 데이터 출력
	}
	return err
}
