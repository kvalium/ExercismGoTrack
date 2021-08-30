package pangram

import (
	"unicode"
)

var alphabetLength = 26

var Exists struct{}

// IsPangram determines whether or not given string is a pangram
func IsPangram(s string) bool {
	return countUniqueLetters(s) == alphabetLength
}

func countUniqueLetters(s string) int {
	uniqueLetters := make(map[rune]struct{})

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		lowerRune := unicode.ToLower(r)

		if _, ok := uniqueLetters[lowerRune]; !ok {
			uniqueLetters[lowerRune] = Exists
		}
	}

	return len(uniqueLetters)
}
