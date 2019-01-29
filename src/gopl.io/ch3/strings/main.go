package main

import "fmt"

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {

	s := "hello, world"
	pref := "hel"
	suf := "world"
	cont := "llo"

	fmt.Printf("Found prefix: %t in string: %s\n", HasPrefix(s, pref), s)
	fmt.Printf("Found suffix: %t in string: %s\n", HasSuffix(s, suf), s)
	fmt.Printf("Contains word: %t in string: %s\n", Contains(s, cont), s)
}
