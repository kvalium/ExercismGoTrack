package diffsquares

import (
	"math"
)

// Difference returns the difference between the square of the sum
// and the sum of the squares of the given terms
func Difference(term int) int {
	return SquareOfSum(term) - SumOfSquares(term)
}

// SquareOfSum returns the square of given terms sum
func SquareOfSum(term int) (result int) {
	for i := 0; i <= term; i++ {
		result += i
	}
	return intPowTwo(result)
}

// SumOfSquares returns the given terms sum square
func SumOfSquares(term int) (result int) {
	for i := 0; i <= term; i++ {
		result += intPowTwo(i)
	}
	return
}

func intPowTwo(term int) int {
	return int(math.Pow(float64(term), 2))
}
