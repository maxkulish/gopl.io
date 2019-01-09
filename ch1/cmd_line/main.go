package main

import (
	"fmt"
	"os"
	"strings"
)

// Long variant with range
func main1()  {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "\n"
	}
	fmt.Println(s)
}

// Shortest realization
func main()  {
	fmt.Println(strings.Join(os.Args, "\n"))
}