package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

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

// randomNonPalindrome returns a string guaranteed not to be a palindrome,
// whose length and contents are derived from the pseudo-random number generator rng.
func randomNonPalindrome(rng *rand.Rand) string {
	minlen := 2
	maxlen := 25

	// any string with length 0 or 1 is a palindrome by definition
	// this adjustment in n makes sure to produce a random length in interval [2, n)
	n := rng.Intn(maxlen-minlen) + minlen

	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
	again:
		r1 := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		r2 := r1
		for r2 == r1 {
			r2 = rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		}

		if !unicode.IsLetter(r1) || !unicode.IsLetter(r2) {
			goto again
		}

		runes[i] = r1
		runes[n-1-i] = r2
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
