package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	// defer 함수 (panic 호출 되면 실행시킬 함수 정의)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("file open error : ", r)
		}
	}()

	f, err := os.Open(filename)
	//	fmt.Println(err)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1 : ", f.Name())
	}

	defer f.Close()
}

func main() {
	fileOpen("undefined.txt")
	fmt.Println("and main")
}
