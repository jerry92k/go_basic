package main

import (
	"fmt"
	oper "section12/arithmetic" // alias oper로 사용. 패키지 중복 또는 약자로 사용
)

func main() {
	/* 사용자 패키지 작성 & Go 문서화
	- 기준은 GOPATH/src
	- 패키지 폴더명(디렉토리명) 명확하게 지정
	- 패키지 파일의 package 이름으로 사용 -> 길면 alias
	- package main을 제외하고 package 문서에 등록
	- 지금까지 우리는 패키지를 사용해 왔다.
	- 기본적으로 GOROOT의 패키지(pkg) 검색 -> GOPATH의 패키지(src/pkg) 검색
	// go install 명령어 실행의 경우 -> GOPATH/pkg에 등록 (문서화)
	// godoc - http:6060 (임의의포트) -> pkg 이동 -> 본인 패키지 매소드 및 주석 확인 (패키지, 타입, 매소드) 주석

	// 패키지 사용 예제(사칙연산)
	*/

	nums := oper.Numbers{100, 10}
	fmt.Println("package used(1) :", nums.Plus())
	fmt.Println("package used(2) :", nums.Minus())
	fmt.Println("package used(3) :", nums.Product())
	fmt.Println("package used(4) :", nums.Devide())
	fmt.Println("package used(5) :", nums.SqaurePlus())
	fmt.Println("package used(6) :", nums.SquareMinus())
}
