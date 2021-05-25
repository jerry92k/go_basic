package main

import "fmt"

func main() {
	i := 100

	if i >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120 미만")
	} else if i >= 80 && i < 100 {
		fmt.Println("80 이상 100 미만")
	} else {
		fmt.Println("그외")
	}

}
