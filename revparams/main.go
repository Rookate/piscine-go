package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[1:]
	for i := len(programName) - 1; i >= 0; i-- {
		arg := programName[i]
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
