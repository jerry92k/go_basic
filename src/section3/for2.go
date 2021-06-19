package main

import "fmt"

func main() {
	// 예제1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1
	}
	fmt.Println(sum1)

	//예제 2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		// j:=i++ // Go에서 후취연산은 반환값 없음
	}
	fmt.Println(sum2)
	sum3 := 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println(sum3)

	//예제 4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {

		// for i, j := 0, 0; i <= 10; i++, j+=10 { // 후취연산 오류
		fmt.Println(i, j)
	}

}
