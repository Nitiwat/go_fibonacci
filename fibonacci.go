package main

import (
	"fmt"
)

func fibonacci() func() int {
	num1 := 0
	num2 := 1
	return func() int {
		next := num1 + num2
		num1 = num2
		num2 = next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
