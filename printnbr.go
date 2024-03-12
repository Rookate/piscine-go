package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	p := 1
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		p = -1
	}
	if n != 0 {
		a := (n / 10) * p
		if a != 0 {
			PrintNbr(a)
		}
		numb := (n % 10 * p) + '0'
		z01.PrintRune(rune(numb))
	}
}
