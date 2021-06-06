package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}
func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	//함수를 배열에 할당
	f := []func(int, int) int{multiply, sum}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println(a, b)

	fmt.Println(f[0](20, 10), f[1](20, 10))

	//함수를 변수에 할당
	var f1 func(int, int) int = multiply
	f2 := sum

	fmt.Println("ex2 : ", f1(10, 10))
	fmt.Println("ex2 : ", f2(10, 10))

	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}

	fmt.Println(m["mul_func"](10, 10))

	fmt.Println(m["sum_func"](20, 10))

}
