package main

import "fmt"

func main() {

	// 배열은 용량, 길이가 항상 같다
	/* 배열 vs 슬라이스
	   ㅇ 배열
	    - 길이 고정
	    - 값 타입
	    - 전달시 복사 전달
	    - 비교연산자 사용 가능
	   ㅇ 슬라이스
	    - 길이가변
	    - 참조 타입
	    - 전달시 참조 값 전달
	    - 비교연산자 사용 불가
	   ㅁ 대부분 슬라이스를 사용한다.

	   ㅇ cap() : 배열, 슬라이스 용량
	   ㅇ len() : 배열, 슬라이스 개수
	*/

	//예제1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} //  선언한 개수와 인자의 개수가 같아야함
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}         // 할당 안된 곳은 0 할당
	arr6 := [...]int{1, 2, 3, 4, 5} // 배열 크기 자동 맞춤
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	arr1[2] = 5

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%d %v\n", len(arr5), arr5)
	fmt.Printf("%d %v\n", len(arr6), arr6)
	fmt.Printf("%d %v\n", len(arr7), arr7)

	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "park"}

	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)
}
