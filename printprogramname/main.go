package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	z := os.Args[0]

	for _, i := range z {
		if i != '/' && i != '.' {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
