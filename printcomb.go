package piscine

// La fonction PrintComb génère et affiche toutes les combinaisons uniques
// de trois chiffres différents de '012' à '789'.
func PrintComb() {
	for x := '0'; x <= '9'; x++ {
		for y := x + 1; y <= '9'; y++ {
			for z := y + 1; z <= '9'; z++ {
				if x < '7' {
					println(string(x) + string(y) + string(z) + ", ")
				} else {
					println(string(x) + string(y) + string(z))
				}
			}
		}
	}
}
