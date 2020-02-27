// Package word provides utilities for word games.
package word

import (
	"math/rand"
	"unicode"
)

// IsPalindrome reports whether s reads the same forward and backward
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	return string(runes)
}
