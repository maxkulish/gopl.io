// Nonempty is an example of in-place slice algorithm
package main

import (
	"fmt"
)

// nonempty returns a slice holding only the non-empty strings
// The underlying array is modified during the call
func nonempty(str []string) []string {
	i := 0
	for _, s := range str {
		if s != " " && s != "" {
			str[i] = s
			i++
		}
	}

	return str[:i]
}

func main()  {
	str := []string{"", "mon", "tue", "wen", " ", "thr"}

	str = nonempty(str)

	fmt.Printf("%q\n", str)
}


