// Package isogram returns isogram status
package isogram

import (
	"unicode"
)

// IsIsogram status for a given string
func IsIsogram(s string) bool {
	cleanedString := formalize(s)
	for _, r := range cleanedString {
		cleanedString = cleanedString[1:]
		if runeinSlice(r, cleanedString) {
			return false
		}
	}
	return true
}

func formalize(s string) []rune {
	var processedString []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			processedString = append(processedString, unicode.ToLower(r))
		}
	}
	return processedString
}

func runeinSlice(s rune, list []rune) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}
