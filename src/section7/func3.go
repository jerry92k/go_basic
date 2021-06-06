package main

import "fmt"

func multiply(x int, y int) (int, int) { //(x,y int 가능)
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
	a, b := multiply(1, 2)
	c, _ := multiply(5, 2)
	_, d := multiply(1, 10)

	fmt.Println("ex1: ", a, b)
	fmt.Println("ex1: ", c, d)

	x1, x2, x3, x4, x5 := arrayMultiply(1, 1, 1, 1, 1)
	x6, _, x7, _, x10 := arrayMultiply(1, 1, 1, 1, 1)

	fmt.Println("ex2: ", x1, x2, x3, x4, x5)
	fmt.Println("ex2: ", x6, x7, x10)
}
