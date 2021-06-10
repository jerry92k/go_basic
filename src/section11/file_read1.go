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
	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의 (오류 메시지 확인)
	// 예외 처리 정말 중요!
	// 탐색 Seek 중요

	// 파일 읽기 예제
	// 파일 열기

	file, err := os.Open("sample.txt")
	errCheck1(err)

	// 읽기 예제1
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 획득
	errCheck2(err)

	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담는다

	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) :", fileInfo, "\n")
	fmt.Println("파일 이름(2)  :", fileInfo.Name(), "\n")
	fmt.Println("파일 크기(3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간(4) :", fileInfo.ModTime(), "\n")
	fmt.Println("읽기 작업(1) 완료 (%d bytes)\n\n", ct1)
	fmt.Println(fd1)         // 바이트 형태 그대로
	fmt.Println(string(fd1)) // 타입 변환한 경우

	fmt.Println("==================================")

	// 읽기 예제2(탐색 : seek(offset))
	o1, err := file.Seek(20, 0) // o: 처음 위치, 1: 현재위치, 2: 끝 offset
	errCheck1(err)
	fmt.Println(o1)
	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck2(err)
	fmt.Println(ct2)
	fmt.Println(string(fd2), "\n")
	fmt.Println("----------------------------------")

	//읽기 예제3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치부터 읽어옴. 요즘은 offset으로 읽으며 file io 하는 것 보다 한번에 다 읽어와서 메모리에서 파싱하는게 더 빠름.

	errCheck2(err)
	fmt.Println(o2)
	fmt.Println(ct3)
	fmt.Println(string(fd3), "\n")
	fmt.Println("----------------------------------")

	defer file.Close()
}
