package main

import "fmt"

func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Remove 1:", remove1(s1, 2))

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Remove 2:", remove2(s2, 2))
}
