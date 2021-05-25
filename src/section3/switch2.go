package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 예제1
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i는 50이상 100미만")
	case i >= 25 && i < 50:
		fmt.Println("i는 25이상 50미만")
	default:
		fmt.Println("i는 25미만")
	}
}
