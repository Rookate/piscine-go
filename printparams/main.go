package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[1:]
	for i := 0; i < len(programName); i++ {
		arg := programName[i]
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
