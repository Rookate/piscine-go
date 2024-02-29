package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, s := range s { // Parcours chaque caractère de la chaîne
		z01.PrintRune(s) // Affiche le caractère
	}
}
