package main

import (
	"fmt"
	"sort"
)

func main()  {

	ages := map[string]int{
		"bob": 42,
		"alice": 24,
		"mikle": 83,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	if age, ok := ages["nestor"]; !ok {
		fmt.Println("Can't find such name:", age)
	}

	newAges := map[string]int{
		"bob": 42,
		"alice": 24,
		"mikle": 83,
	}

	// compare two maps
	fmt.Printf("Two maps are equal: %t\n", equal(ages, newAges))

	// True if equal is written incorrectly
	fmt.Printf("Two maps are equal: %t\n", equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}