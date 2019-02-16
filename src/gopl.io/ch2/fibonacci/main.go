package main

import "fmt"

func fibonacci(n int) int {

	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, y+x
	}
	return x
}

func main() {
	n := 13
	fmt.Printf("Number: %d\nResult: %d\n", n, fibonacci(n))
}
