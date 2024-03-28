package main

import (
	"os"

	"github.com/01-edu/z01"
)

func CanWrite(first, second string) bool {
	indexFirst := 0

	for _, ch := range second {
		if indexFirst < len(first) && byte(ch) == first[indexFirst] {
			indexFirst++
		}
	}
	return indexFirst == len(first)
}

func trimQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		// On fait len(s) - 1 car on récupère l'index du dernier caractère notre chaine. En go les indices de chaines de caractères commence à 0
		return s[1 : len(s)-1]
	}
	return s
}

func main() {
	if len(os.Args) == 3 {
		firstString := trimQuotes(os.Args[1])  // Retirer les guillemets du premier argument si présents
		secondString := trimQuotes(os.Args[2]) // Retirer les guillemets du deuxième argument si présents
		if CanWrite(firstString, secondString) {
			for _, ch := range firstString {
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
		}
	}
}
