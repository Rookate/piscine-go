package piscine

import (
	"github.com/01-edu/z01"
)

func IterativeFactorial(nb int) int {
	z01.PrintRune(' ')

	result := 1

	for i := 1; i <= nb; i++ {
		result = result * i
	}

	return result
}
