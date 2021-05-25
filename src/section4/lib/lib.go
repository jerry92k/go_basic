// 라이브러리 접근제어(1)

package lib

import "fmt"

func CheckNum(c int32) bool {
	return c > 10
}

func init() {
	fmt.Println("lib init")
}
