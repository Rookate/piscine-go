package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for x := '0'; x <= '9'; x++ { // Boucle pour le premier chiffre
		for y := '0'; y <= '9'; y++ { // Boucle pour le deuxième chiffre
			for z := '0'; z <= '9'; z++ { // Boucle pour le troisième chiffre
				for w := '0'; w <= '9'; w++ { // Boucle pour le quatrième chiffre
					if x < z || (x == z && y < w) { // Vérifie si la combinaison est dans l'ordre croissant
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(z)
						z01.PrintRune(w)
						if !(x == '9' && y == '8' && z == '9' && w == '9') { // Vérifie si ce n'est pas la dernière combinaison
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n') // Saut de ligne à la fin
}
