// basename removes directory components and a .suffix
// home/dev/file.go => file
package main

import "fmt"

func basename(s string) string {

	// remove lst '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	s := "/Users/mk/Code/ogorod/celeryconfig.py"

	fmt.Println(basename(s))
}
