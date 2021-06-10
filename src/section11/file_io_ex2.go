package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기, 버퍼사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, Wrtier, Read 매소드 사용 가능

	/*
			  type Reader interface{
			    Read(p []byte)(n int, err error)
			  }

			  type Writer interface{
			    Read(p []byte)(n int, err error)
			  }

			  //즉, bufio의 NewReader, NewWriter를 통해서 객체 반환후 매소드 사용

			  // bufio 패키지( buffered io) 패키지
			  // https://golang.org/pkg/bufio의

			  // 파일 열기
			  // 두번째 매개변수 확인
			  // https://golang.org/pkg/os/#pkg-constants

		    상태
		    a --->a
		    b --->ab
		    c --->abc
		    d --->abcd
		    e --->e    => abcd 배출

		    // bufio 파일 쓰기 예제
	*/

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777)) //O_CREATE : 없으면 만들고, O_RDWR : 있으면 읽고 쓰기 가능하게
	errCheck(err)

	wt := bufio.NewWriter(file) // Write 반환(버퍼 사용)

	wt.WriteString("Hello golang! \n File wirte test1\n")
	wt.Write([]byte("Hello golang! \n File wirte test2\n"))

	errCheck(err)

	fmt.Printf("사용한 buffer size (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 buffer size (%d bytes)\n", wt.Available())
	fmt.Printf("전체 buffer size (%d bytes)\n", wt.Size())

	wt.Flush() // 버퍼에 있는 내용을 디스크에 기록

	fmt.Println("쓰기작업 완료")

	rt := bufio.NewReader(file) // Reader 바환
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 :", fi)
	fmt.Println("파일 이름 :", fi.Name())
	fmt.Println("파일 크기 :", fi.Size())
	fmt.Println("파일 수정시간 :", fi.ModTime())

	fmt.Println("=====================")

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b)
	// rt.Read(b) // 또 호출하면 위치가 바뀜

	fmt.Printf("전체 buffer size : (%d) bytes\n", rt.Size())

	fmt.Printf("읽기 작업 완료! : %d  bytes \n", data)

	fmt.Println("=====================")

	fmt.Println(string(b))

}
