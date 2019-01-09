package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Long variant with range
func echo_long()  {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "\n"
	}
	fmt.Println(s)
}

// Shortest realization
func echo_short()  {
	fmt.Println(strings.Join(os.Args, "\n"))
}

func main()  {
	start := time.Now()
	echo_long()
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	echo_short()
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())

}