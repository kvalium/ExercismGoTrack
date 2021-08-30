// Package romannumerals convert an arabic number to the corresponding
// roman number.
package romannumerals

import (
	"fmt"
	"strings"
)

var romanToArabic = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

// ToRomanNumeral takes an arabic number and returns
// its corresponding roman translation.
func ToRomanNumeral(n int) (string, error) {
	var r string
	if n > 3000 || n <= 0 {
		return r, fmt.Errorf("cannot convert null or negative number or number over 3000")
	}
	for roman, arabic := range romanToArabic {
		occ := n / arabic
		if occ == 0 {
			continue
		}
		r += strings.Repeat(roman, occ)
		n -= occ * arabic
	}
	return r, nil
}
