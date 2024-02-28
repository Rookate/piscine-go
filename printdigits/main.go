package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var aRune rune
	for aRune = 48; aRune <= 57; aRune++ {
		z01.PrintRune(aRune)
	}
	z01.PrintRune('\n')
}
