package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func min(vals ...int) int {
	min := vals[0]
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}

	return min
}

func max(vals ...int) int {
	max := vals[0]
	for _, val := range vals[1:] {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	fmt.Println("total:", sum(1, 3, 4, 5, 6))
	fmt.Println("total:", sum(3))

	values := []int{10, 2, 3, 4, -2, 13}
	fmt.Println("Sum:", sum(values...))

	fmt.Println("Min:", min(values...))

	fmt.Println("Max:", max(values...))
}
