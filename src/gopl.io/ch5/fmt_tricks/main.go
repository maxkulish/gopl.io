package main

import "fmt"

func main() {
	var depth int

	for i := 0; i < 10; i++ {
		depth += 2
		fmt.Printf("%*s padding: %d\n", depth*1, "=>", depth)
	}
}
