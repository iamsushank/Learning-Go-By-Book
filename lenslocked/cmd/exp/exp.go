package main

import (
	"fmt"
)

func main() {
	Demo()
	Demo(1)
	Demo(1, 2, 3)
}
func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
