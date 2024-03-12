package piscine

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for _, ch := range a {
		f(ch)
	}
	z01.PrintRune('\n')
}
